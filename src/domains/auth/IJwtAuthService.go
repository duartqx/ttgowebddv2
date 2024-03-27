package auth

import (
	"net/http"
	"time"

	u "github.com/duartqx/ddgobase/src/domains/user"
)

type IJwtAuthService interface {
	ValidateAuth(authorization string, cookie *http.Cookie) (ISessionUser, error)
	Login(user u.IUser) (token string, expiresAt time.Time, err error)
	Logout(authorization string, cookie *http.Cookie) error
}
