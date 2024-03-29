package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/duartqx/ttgowebddv2/internal/api/dto"
	h "github.com/duartqx/ttgowebddv2/internal/api/http"
	e "github.com/duartqx/ttgowebddv2/internal/common/errors"
	a "github.com/duartqx/ttgowebddv2/internal/domains/auth"
	u "github.com/duartqx/ttgowebddv2/internal/domains/user"
)

type JwtController struct {
	jwtService a.IJwtAuthService
}

var jwtController *JwtController

func GetJwtController(jwtService a.IJwtAuthService) *JwtController {
	if jwtController == nil {
		jwtController = &JwtController{
			jwtService: jwtService,
		}
	}
	return jwtController
}

func (jc JwtController) Login(w http.ResponseWriter, r *http.Request) {

	var userDTO dto.UserLogin

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := u.GetNewUser().SetEmail(userDTO.Email).SetPassword(userDTO.Password)

	token, expiresAt, err := jc.jwtService.Login(user)

	if err != nil {
		h.ErrorResponse(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     h.AuthCookieName,
		Value:    token,
		Expires:  expiresAt,
		Secure:   true,
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(
		dto.LoginResponse{
			Token:     token,
			ExpiresAt: expiresAt,
			Status:    true,
		},
	); err != nil {
		h.ErrorResponse(w, err)
	}
}

func (jc JwtController) Logout(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie(h.AuthCookieName)

	err := jc.jwtService.Logout(r.Header.Get("Authorization"), cookie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	http.SetCookie(w, h.GetInvalidCookie())
}

func (jc JwtController) AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie(h.AuthCookieName)
		claimsUser, err := jc.jwtService.ValidateAuth(
			r.Header.Get("Authorization"), cookie,
		)

		if err != nil {
			http.SetCookie(w, h.GetInvalidCookie())
			h.ErrorResponse(w, err)
			return
		}

		// Injects the User Information into the request context
		ctx := context.WithValue(r.Context(), "user", claimsUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (jc JwtController) UnauthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie(h.AuthCookieName)
		claimsUser, _ := jc.jwtService.ValidateAuth(
			r.Header.Get("Authorization"), cookie,
		)

		if claimsUser != nil {
			http.SetCookie(w, h.GetInvalidCookie())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
