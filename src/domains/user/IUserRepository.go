package user

type IUserRepository interface {
	FindById(user IUser) error

	FindByEmail(user IUser) error
	ExistsByEmail(email string) bool

	Create(user IUser) error
	Update(user IUser) error
	Delete(user IUser) error
}
