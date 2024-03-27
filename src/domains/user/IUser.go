package user

type IUser interface {
	GetId() int
	SetId(id int) *User

	GetName() string
	SetName(name string) *User

	GetPassword() string
	SetPassword(password string) *User

	GetEmail() string
	SetEmail(email string) *User
}
