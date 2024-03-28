package task

import (
	"time"

	u "github.com/duartqx/ttgowebddv2/src/domains/user"
)

type Task struct {
	Id          int        `json:"id" db:"id"`
	Tag         string     `json:"tag" db:"tag"`
	Sprint      string     `json:"sprint" db:"sprint"`
	Description string     `json:"description" db:"description"`
	Completed   bool       `json:"completed" db:"completed"`
	StartAt     *time.Time `json:"start_at" db:"start_at"`
	EndAt       *time.Time `json:"end_at" db:"end_at"`

	UserId int     `json:"-" db:"user_id"`
	User   *u.User `json:"user" db:"user"`
}

func GetNewTask() *Task {
	return &Task{User: u.GetNewUser()}
}

func (t Task) GetId() int {
	return t.Id
}

func (t *Task) SetId(id int) *Task {
	t.Id = id
	return t
}

func (t Task) GetTag() string {
	return t.Tag
}

func (t *Task) SetTag(tag string) *Task {
	t.Tag = tag
	return t
}

func (t Task) GetSprint() string {
	return t.Sprint
}

func (t *Task) SetSprint(sprint string) *Task {
	t.Sprint = sprint
	return t
}

func (t Task) GetDescription() string {
	return t.Description
}

func (t *Task) SetDescription(description string) *Task {
	t.Description = description
	return t
}

func (t Task) GetCompleted() bool {
	return t.Completed
}

func (t *Task) SetCompleted(completed bool) *Task {
	t.Completed = completed
	return t
}

func (t Task) GetStartAt() *time.Time {
	return t.StartAt
}

func (t *Task) SetStartAt(startAt *time.Time) *Task {
	t.StartAt = startAt
	return t
}

func (t Task) GetEndAt() *time.Time {
	return t.EndAt
}

func (t *Task) SetEndAt(endAt *time.Time) *Task {
	t.EndAt = endAt
	return t
}

func (t Task) GetUserId() int {
	return t.UserId
}

func (t *Task) SetUserId(id int) *Task {
	t.UserId = id
	return t
}

func (t Task) GetUser() u.IUser {
	return t.User
}

func (t *Task) SetUser(user u.IUser) *Task {
	t.User.
		SetId(user.GetId()).
		SetName(user.GetName()).
		SetEmail(user.GetEmail())

	t.UserId = user.GetId()

	return t
}

func (t *Task) ToLocaltime() *Task {
	if t.GetStartAt() != nil {
		localtimeStartAt := t.GetStartAt().Local()
		t.SetStartAt(&localtimeStartAt)
	}
	if t.GetEndAt() != nil {
		localtimeEndAt := t.GetEndAt().Local()
		t.SetEndAt(&localtimeEndAt)
	}
	return t
}
