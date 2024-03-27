package auth

import (
	u "github.com/duartqx/ddgobase/src/domains/user"
)

type ISessionUser interface {
	GetId() int
	GetEmail() string
	GetName() string

	SetId(id int) ISessionUser
	SetEmail(email string) ISessionUser
	SetName(name string) ISessionUser

	SetFromAnother(user ISessionUser) ISessionUser
	ToUser() u.IUser
}
