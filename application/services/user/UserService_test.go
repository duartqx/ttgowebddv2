package user_test

import (
	"testing"

	s "github.com/duartqx/ddgobase/application/services/user"
	u "github.com/duartqx/ddgobase/domains/user"
	r "github.com/duartqx/ddgobase/infrastructure/repository"
)

var userService u.IUserService = s.GetUserService(r.GetMockUserRepository())

func TestCreate(t *testing.T) {
	tests := []struct {
		name string
		user u.User
		err  bool
	}{
		{
			name: "FailEmailInvalid",
			user: u.User{Name: "Test User 1", Email: "", Password: "randompassword"},
			err:  true,
		},
		{
			name: "FailEmailExists",
			user: u.User{Name: "Test User 1", Email: "test1@test1.com", Password: "randompassword"},
			err:  true,
		},
		{
			name: "FailNameInvalid",
			user: u.User{Name: "", Email: "teste7@teste7.com", Password: "randompassword"},
			err:  true,
		},
		{
			name: "FailNameEmailInvalid",
			user: u.User{Name: "", Email: "", Password: "randompassword"},
			err:  true,
		},
		{
			name: "FailInvalidPassword",
			user: u.User{Name: "Test User 1", Email: "teste99@99.com", Password: ""},
			err:  true,
		},
		{
			name: "PassEmailDoesNotExists",
			user: u.User{Name: "Test User 6", Email: "test6@test6.com", Password: "randompassword"},
			err:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := userService.Create(&tt.user)

			t.Logf("Create Error: %v", err)

			if tt.err && err == nil {
				t.Fatalf("Expected an error, got nil")
			} else if !tt.err && err != nil {
				t.Fatalf("%s: Expected nil, got an error", err.Error())
			}

			if err == nil && tt.user.GetId() == 0 {
				t.Fatalf("Expected user id to update")
			}
		})
	}
}

func TestFindById(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		user     u.IUser
		err      bool
	}{
		{name: "PassIdExists1", expected: 1, user: &u.User{Id: 1}, err: false},
		{name: "PassIdExists2", expected: 2, user: &u.User{Id: 2}, err: false},
		{name: "PassIdExists3", expected: 3, user: &u.User{Id: 3}, err: false},
		{name: "FailIdDoesNotExists1", expected: 99, user: &u.User{Id: 99}, err: true},
		{name: "FailIdDoesNotExists2", expected: 999, user: &u.User{Id: 999}, err: true},
		{name: "FailIdDoesNotExists3", expected: 333, user: &u.User{Id: 333}, err: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := userService.FindById(tt.user)

			t.Logf("FindById Error: %v", err)

			if tt.err && err == nil {
				t.Fatalf("Expected an error with id == %d, got nil", tt.expected)
			} else if !tt.err && err != nil {
				t.Fatalf(
					"%s: Expected no error with id == %d, got error",
					err.Error(),
					tt.expected,
				)
			}

			if err == nil && tt.user.GetId() != tt.expected {
				t.Fatalf(
					"Expected user id to match the one on test case, got %d, expected %d",
					tt.user.GetId(),
					tt.expected,
				)
			}
		})
	}
}

func TestUpdatePassword(t *testing.T) {
	tests := []struct {
		name        string
		user        u.User
		oldPassword string
		err         bool
	}{
		{
			name: "PassNewPassword",
			user: u.User{
				Id:       1,
				Name:     "Test User 1",
				Email:    "test1@test1.com",
				Password: "New Password",
			},
			oldPassword: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
			err:         false,
		},
		{
			name: "FailInvalidPassword",
			user: u.User{
				Id:       1,
				Name:     "Test User 1",
				Email:    "test1@test1.com",
				Password: "",
			},
			oldPassword: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
			err:         true,
		},
		{
			name: "FailInvalidId",
			user: u.User{
				Id:       0,
				Name:     "Test User 1",
				Email:    "test1@test1.com",
				Password: "randompassword",
			},
			oldPassword: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
			err:         true,
		},
		{
			name: "FailInvalidIdInvalidPassword",
			user: u.User{
				Id:       0,
				Name:     "Test User 1",
				Email:    "test1@test1.com",
				Password: "",
			},
			oldPassword: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
			err:         true,
		},
		{
			name: "PassSamePassword",
			user: u.User{
				Id:       1,
				Name:     "Test User 1",
				Email:    "test1@test1.com",
				Password: "randompassword",
			},
			oldPassword: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
			err:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := userService.UpdatePassword(&tt.user)

			t.Logf("UpdatePassword Error: %v", err)

			if tt.err && err == nil {
				t.Fatalf("%s: Expected an error, got nil", err.Error())
			} else if !tt.err && err != nil {
				t.Fatalf("%s: Expected nil, got error", err.Error())
			}

			if err == nil {
				if tt.user.GetPassword() == "" {
					t.Fatalf("%s: Password change is invalid", tt.user.GetPassword())
				}

				if tt.user.GetPassword() == tt.oldPassword {
					t.Fatalf("%s: Password did not changed", tt.user.GetPassword())
				}
			}
		})
	}
}
