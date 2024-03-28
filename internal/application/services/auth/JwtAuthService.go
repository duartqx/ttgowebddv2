package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	e "github.com/duartqx/ttgowebddv2/internal/common/errors"
	a "github.com/duartqx/ttgowebddv2/internal/domains/auth"
	u "github.com/duartqx/ttgowebddv2/internal/domains/user"
)

type JwtAuthService struct {
	userRepository    u.IUserRepository
	sessionRepository a.ISessionRepository
	secret            *[]byte
}

var jwtAuthService *JwtAuthService

func GetJwtAuthService(
	userRepository u.IUserRepository,
	sessionRepository a.ISessionRepository,
	secret *[]byte,
) *JwtAuthService {
	if jwtAuthService == nil {
		jwtAuthService = &JwtAuthService{
			userRepository:    userRepository,
			sessionRepository: sessionRepository,
			secret:            secret,
		}
	}
	return jwtAuthService
}

func (jas JwtAuthService) keyFunc(t *jwt.Token) (interface{}, error) {
	return *jas.secret, nil
}

func (jas JwtAuthService) generateToken(user a.ISessionUser, expiresAt time.Time) (string, error) {

	tokenStr, _ := jwt.NewWithClaims(
		jwt.SigningMethodHS256, a.GetNewPopulatedClaims(user, expiresAt),
	).SignedString(*jas.secret)

	return tokenStr, nil
}

func (jas JwtAuthService) getUnparsedToken(auth string, c *http.Cookie) string {
	if auth != "" {
		token, found := strings.CutPrefix(auth, "Bearer ")
		if found {
			return token
		}
	}
	if c != nil {
		return c.Value
	}
	return ""
}

func (jas JwtAuthService) ValidateAuth(
	authorization string, cookie *http.Cookie,
) (a.ISessionUser, error) {

	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return nil, fmt.Errorf("%w: Missing Token", e.Unauthorized)
	}

	claims := a.GetNewClaims()

	parsedToken, err := jwt.ParseWithClaims(unparsedToken, claims, jas.keyFunc)
	if err != nil || !parsedToken.Valid {
		if jas.sessionRepository != nil {
			go jas.sessionRepository.Delete(claims.ISessionUser)
		}
		return nil, fmt.Errorf("%w: Expired session", e.Unauthorized)
	}

	return claims.ISessionUser, nil
}

func (jas JwtAuthService) Login(user u.IUser) (token string, expiresAt time.Time, err error) {

	if user.GetEmail() == "" || user.GetPassword() == "" {
		return token, expiresAt, fmt.Errorf("%w: Invalid Email or Password", e.BadRequestError)
	}

	dbUser := u.GetNewUser().SetEmail(user.GetEmail())

	if err := jas.userRepository.FindByEmail(dbUser); err != nil {
		return token, expiresAt, fmt.Errorf("%w: Invalid Email", e.Unauthorized)
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(dbUser.GetPassword()), []byte(user.GetPassword()),
	); err != nil {
		return token, expiresAt, fmt.Errorf("%w: Invalid Password", e.BadRequestError)
	}

	createdAt := time.Now()
	expiresAt = createdAt.Add(time.Hour * 12)

	claimsUser := a.GetNewSessionUser()

	claimsUser.
		SetId(dbUser.GetId()).
		SetEmail(dbUser.GetEmail()).
		SetName(dbUser.GetName())

	token, err = jas.generateToken(claimsUser, expiresAt)
	if err != nil {
		return "", expiresAt, fmt.Errorf("%w: Could not generate token", e.InternalError)
	}

	if jas.sessionRepository != nil {
		jas.sessionRepository.Set(claimsUser, createdAt)
	}

	return token, expiresAt, nil
}

func (jas JwtAuthService) Logout(authorization string, cookie *http.Cookie) error {
	unparsedToken := jas.getUnparsedToken(authorization, cookie)
	if unparsedToken == "" {
		return fmt.Errorf("%w: Missing Token", e.Unauthorized)
	}

	claims := a.GetNewClaims()

	if _, err := jwt.ParseWithClaims(unparsedToken, claims, jas.keyFunc); err == nil {
		if jas.sessionRepository != nil {
			go jas.sessionRepository.Delete(claims.ISessionUser)
		}
	}

	return nil
}
