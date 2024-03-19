package user

type IUserService interface {
	FindById(user IUser) error

	Create(user IUser) error
	Update(user IUser) error
	Delete(user IUser) error

	UpdatePassword(user IUser) error
}
