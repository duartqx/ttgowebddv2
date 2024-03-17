package user

type IUserService interface {
	FindById(id int) (IUser, error)

	Create(user IUser) error
	Update(user IUser) error
	Delete(user IUser) error

	UpdatePassword(user IUser) error
}
