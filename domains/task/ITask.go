package task

import (
	"time"

	u "github.com/duartqx/ttgowebddv2/domains/user"
)

type ITask interface {
	GetId() int
	SetId(id int) ITask

	GetTag() string
	SetTag(tag string) ITask

	GetSprint() string
	SetSprint(sprint string) ITask

	GetDescription() string
	SetDescription(description string) ITask

	GetCompleted() bool
	SetCompleted(completed bool) ITask

	GetStartAt() *time.Time
	SetStartAt(startAt *time.Time) ITask

	GetEndAt() *time.Time
	SetEndAt(endAt *time.Time) ITask

	GetUserId() int
	SetUserId(id int) ITask

	GetUser() u.IUser
	SetUser(user u.IUser) ITask

	ToLocaltime() ITask
}
