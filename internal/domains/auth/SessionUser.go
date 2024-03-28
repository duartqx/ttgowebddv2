package auth

import (
	"github.com/duartqx/ddgobase/internal/domains/user"
)

type SessionUser struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GetNewSessionUser() *SessionUser {
	return &SessionUser{}
}

func (u SessionUser) GetId() int {
	return u.Id
}

func (u *SessionUser) SetId(id int) ISessionUser {
	u.Id = id
	return u
}

func (u SessionUser) GetEmail() string {
	return u.Email
}

func (u *SessionUser) SetEmail(email string) ISessionUser {
	u.Email = email
	return u
}

func (u SessionUser) GetName() string {
	return u.Name
}

func (u *SessionUser) SetName(name string) ISessionUser {
	u.Name = name
	return u
}

func (u *SessionUser) SetFromAnother(user ISessionUser) ISessionUser {
	u.SetId(user.GetId()).SetEmail(user.GetEmail()).SetName(user.GetName())
	return u
}

func (u *SessionUser) ToUser() user.IUser {
	return user.GetNewUser().SetId(u.Id).SetName(u.Name).SetEmail(u.Email)
}
