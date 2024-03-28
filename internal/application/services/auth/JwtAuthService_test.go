package auth_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"

	j "github.com/duartqx/ttgowebddv2/internal/application/services/auth"
	u "github.com/duartqx/ttgowebddv2/internal/domains/user"
	s "github.com/duartqx/ttgowebddv2/internal/infrastructure/repository"
	r "github.com/duartqx/ttgowebddv2/internal/infrastructure/repository/sqlite"
)

var (
	db *sqlx.DB

	secret            = []byte("secret")
	sessionRepository *s.SessionRepository
	userRepository    *r.UserRepository
	jwtAuthService    *j.JwtAuthService
)

func TestMain(m *testing.M) {
	db = r.GetInMemoryDB("jwtauthservice")
	defer db.Close()

	r.Seed(db)

	sessionRepository = s.GetSessionRepository()
	userRepository = r.GetUserRepository(db)
	jwtAuthService = j.GetJwtAuthService(
		userRepository, sessionRepository, &secret,
	)

	code := m.Run()

	os.Exit(code)
}

func TestLoginAndValidateAuth(t *testing.T) {
	tests := []struct {
		name string
		user u.User
		err  bool
	}{
		{
			name: "FailEmailInvalid",
			user: u.User{Email: "", Password: "randompassword"},
			err:  true,
		},
		{
			name: "FailWrongPassword",
			user: u.User{Email: "test1@test1.com", Password: "wrongpassword"},
			err:  true,
		},
		{
			name: "FailInvalidPassword",
			user: u.User{Email: "test1@test1.com", Password: ""},
			err:  true,
		},
		{
			name: "FailUserDoesNotExists",
			user: u.User{Email: "teste99@teste99.com", Password: "randompassword"},
			err:  true,
		},
		{
			name: "PassCorrectTest1",
			user: u.User{Email: "test1@test1.com", Password: "randompassword"},
			err:  false,
		},
		{
			name: "PassCorrectTest2",
			user: u.User{Email: "test2@test2.com", Password: "randompassword"},
			err:  false,
		},
		{
			name: "PassCorrectTest3",
			user: u.User{Email: "test3@test3.com", Password: "randompassword"},
			err:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Login
			token, expires, err := jwtAuthService.Login(&tt.user)

			t.Logf("Expects Error: %v, Did Error: %v", tt.err, err != nil)

			if tt.err && err == nil {
				t.Fatalf("Expected an error, got nil")
			} else if !tt.err && err != nil {
				t.Fatalf("%s: Expected nil, got an error", err.Error())
			}

			if !tt.err {
				// Validate Expires and ValidateAuth

				t.Logf("Token: %v, Expires: %v, Error: %v", token, expires, err)

				if time.Now().After(expires) {
					t.Fatalf("Expected expiration date to be after time.Now, got %v", expires)
				}

				sessionUser, err := jwtAuthService.ValidateAuth(
					fmt.Sprintf("Bearer %s", token),
					&http.Cookie{Value: token, Expires: expires},
				)

				if err != nil {
					t.Fatalf("%s: Error validating token", err.Error())
				}

				if sessionUser.GetEmail() != tt.user.GetEmail() {
					t.Fatalf(
						"Authenticated user does not match email, got %s, expected %s",
						sessionUser.GetEmail(),
						tt.user.GetEmail(),
					)
				}
			}
		})
	}
}
