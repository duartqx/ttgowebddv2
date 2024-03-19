package mock

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/duartqx/ddgobase/common/errors"
	u "github.com/duartqx/ddgobase/domains/user"
)

var (
	// Password: randompassword
	users []u.IUser = []u.IUser{
		&u.User{
			Id:       1,
			Name:     "Test User 1",
			Email:    "test1@test1.com",
			Password: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
		},
		&u.User{
			Id:       2,
			Name:     "Test User 2",
			Email:    "test2@test2.com",
			Password: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
		},
		&u.User{
			Id:       3,
			Name:     "Test User 3",
			Email:    "test3@test3.com",
			Password: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
		},
		&u.User{
			Id:       4,
			Name:     "Test User 4",
			Email:    "test4@test4.com",
			Password: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
		},
		&u.User{
			Id:       5,
			Name:     "Test User 5",
			Email:    "test5@test5.com",
			Password: "$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW",
		},
	}
	nextId int = len(users)
)

type MockUserRepository struct{}

func GetMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (mur MockUserRepository) FindById(user u.IUser) error {
	userIndex, found := slices.BinarySearchFunc(
		users, user.GetId(), func(dbUser u.IUser, id int) int {
			return cmp.Compare(dbUser.GetId(), id)
		},
	)
	if found {
		dbUser := users[userIndex]
		user.SetName(dbUser.GetName()).
			SetEmail(dbUser.GetEmail()).
			SetPassword(dbUser.GetPassword())
		return nil
	}
	return fmt.Errorf("%w: User not found", errors.NotFoundError)
}

func (mur MockUserRepository) FindByEmail(user u.IUser) error {
	for _, dbUser := range users {
		if dbUser.GetEmail() == user.GetEmail() {
			user.SetPassword(dbUser.GetPassword()).SetName(dbUser.GetName())
			return nil
		}
	}
	return fmt.Errorf("%w: User not found", errors.NotFoundError)
}

func (mur MockUserRepository) ExistsByEmail(email string) bool {
	for _, user := range users {
		if user.GetEmail() == email {
			return true
		}
	}
	return false
}

func (mur MockUserRepository) Create(user u.IUser) error {
	if user.GetEmail() == "" || user.GetName() == "" || user.GetPassword() == "" {
		return fmt.Errorf("%w: Invalid User", errors.BadRequestError)
	}
	if mur.ExistsByEmail(user.GetEmail()) {
		return fmt.Errorf("%w: Invalid Email", errors.BadRequestError)
	}
	user.SetId(nextId)
	nextId = len(users)
	users = append(users, user)
	return nil
}

func (mur MockUserRepository) Update(user u.IUser) error {
	if user.GetId() == 0 || user.GetEmail() == "" || user.GetName() == "" || user.GetPassword() == "" {
		return fmt.Errorf("%w: Invalid User", errors.BadRequestError)
	}
	userIndex, found := slices.BinarySearchFunc(
		users, user.GetId(), func(u u.IUser, id int) int {
			return cmp.Compare(u.GetId(), id)
		},
	)
	if !found {
		return fmt.Errorf("%w: Invalid User", errors.BadRequestError)
	}
	users[userIndex] = user

	return nil
}

func (mur MockUserRepository) Delete(user u.IUser) error {
	var filteredUsers []u.IUser
	for _, u := range users {
		if u.GetId() != user.GetId() {
			filteredUsers = append(filteredUsers, user)
		}
	}
	users = filteredUsers
	return nil
}
