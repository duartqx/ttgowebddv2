package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	v "github.com/duartqx/ddgobase/application/validation"
	e "github.com/duartqx/ddgobase/common/errors"
	u "github.com/duartqx/ddgobase/domains/user"
)

type UserService struct {
	userRepository u.IUserRepository
	validator      *v.Validator
}

var userService *UserService

func GetUserService(userRepository u.IUserRepository) *UserService {
	if userService == nil {
		userService = &UserService{
			userRepository: userRepository,
			validator:      v.NewValidator(),
		}
	}
	return userService
}

func (us UserService) Create(user u.IUser) error {

	if us.userRepository.ExistsByEmail(user.GetEmail()) {
		return fmt.Errorf("%w: Invalid Email", e.BadRequestError)
	}

	if err := us.validator.ValidateStruct(user); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.GetPassword()), 10,
	)
	if err != nil {
		return fmt.Errorf("%w: Invalid Password", e.BadRequestError)
	}

	user.SetPassword(string(hashedPassword))

	if err := us.userRepository.Create(user); err != nil {
		return fmt.Errorf("%w: Internal Error trying to create user", e.InternalError)
	}

	return nil
}

func (us UserService) UpdatePassword(user u.IUser) error {
	if user.GetId() == 0 {
		return fmt.Errorf("%w: Invalid User", e.BadRequestError)
	}

	v := struct {
		Password string `validate:"required,min=8,max=200"`
	}{
		Password: user.GetPassword(),
	}

	if err := us.validator.ValidateStruct(v); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 10)
	if err != nil {
		return fmt.Errorf("%w: Invalid Password", e.BadRequestError)
	}

	user.SetPassword(string(hashedPassword))

	if err := us.userRepository.Update(user); err != nil {
		return fmt.Errorf(
			"%w: Internal Error trying to update the password",
			e.InternalError,
		)
	}

	return nil
}

func (us UserService) Update(user u.IUser) error {
	return us.UpdatePassword(user)
}

func (us UserService) FindById(id int) (u.IUser, error) {
	return us.userRepository.FindById(id)
}

func (us UserService) Delete(user u.IUser) error {
	if user.GetId() == 0 {
		return fmt.Errorf("%w: Invalid User", e.BadRequestError)
	}
	return us.userRepository.Delete(user)
}
