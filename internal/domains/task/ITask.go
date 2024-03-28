package task

import (
	"time"

	u "github.com/duartqx/ttgowebddv2/internal/domains/user"
)

type ITask interface {
	GetId() int
	SetId(id int) *Task

	GetTag() string
	SetTag(tag string) *Task

	GetSprint() string
	SetSprint(sprint string) *Task

	GetDescription() string
	SetDescription(description string) *Task

	GetCompleted() bool
	SetCompleted(completed bool) *Task

	GetStartAt() *time.Time
	SetStartAt(startAt *time.Time) *Task

	GetEndAt() *time.Time
	SetEndAt(endAt *time.Time) *Task

	GetUserId() int
	SetUserId(id int) *Task

	GetUser() u.IUser
	SetUser(user u.IUser) *Task

	ToLocaltime() *Task
}
