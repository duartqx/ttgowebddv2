package user

type IUserService interface {
	FindById(user *User) error

	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error

	UpdatePassword(user *User) error
}
