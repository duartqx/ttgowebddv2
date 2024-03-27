package user

type IUserRepository interface {
	FindById(user *User) error

	FindByEmail(user *User) error
	ExistsByEmail(email string) bool

	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
}
