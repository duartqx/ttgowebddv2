package task

import (
	"time"

	u "github.com/duartqx/ttgowebddv2/domains/user"
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

func (t Task) GetId() int {
	return t.Id
}

func (t *Task) SetId(id int) ITask {
	t.Id = id
	return t
}

func (t Task) GetTag() string {
	return t.Tag
}

func (t *Task) SetTag(tag string) ITask {
	t.Tag = tag
	return t
}

func (t Task) GetSprint() string {
	return t.Sprint
}

func (t *Task) SetSprint(sprint string) ITask {
	t.Sprint = sprint
	return t
}

func (t Task) GetDescription() string {
	return t.Description
}

func (t *Task) SetDescription(description string) ITask {
	t.Description = description
	return t
}

func (t Task) GetCompleted() bool {
	return t.Completed
}

func (t *Task) SetCompleted() ITask {
	t.Completed = !t.Completed
	return t
}

func (t Task) GetStartAt() *time.Time {
	return t.StartAt
}

func (t *Task) SetStartAt(startAt *time.Time) ITask {
	t.StartAt = startAt
	return t
}

func (t Task) GetEndAt() *time.Time {
	return t.EndAt
}

func (t *Task) SetEndAt(endAt *time.Time) ITask {
	t.EndAt = endAt
	return t
}

func (t Task) GetUserId() int {
	return t.UserId
}

func (t *Task) SetUserId(id int) ITask {
	t.UserId = id
	return t
}

func (t Task) GetUser() u.IUser {
	return t.User
}

func (t *Task) SetUser(user u.IUser) ITask {
	t.User.
		SetId(user.GetId()).
		SetName(user.GetName()).
		SetEmail(user.GetEmail())

	t.UserId = user.GetId()

	return t
}

func (t *Task) ToLocaltime() ITask {
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
