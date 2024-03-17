package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ISessionUser
	jwt.RegisteredClaims
}

func GetNewClaims() *CustomClaims {
	return &CustomClaims{
		ISessionUser:     GetNewSessionUser(),
		RegisteredClaims: jwt.RegisteredClaims{},
	}
}

func GetNewPopulatedClaims(user ISessionUser, expiresAt time.Time) *CustomClaims {
	return &CustomClaims{
		ISessionUser: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
}
