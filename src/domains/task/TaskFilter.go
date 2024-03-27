package task

import (
	"time"
)

type completedStatus struct {
	Ignored     int
	Completed   int
	Incompleted int
}

var CompletedStatus = completedStatus{
	Incompleted: 0,
	Completed:   1,
	Ignored:     2,
}

type TaskFilter struct {
	Tag       string    `json:"tag"`
	Completed int       `json:"completed"`
	Sprints   []int     `json:"sprints"`
	StartAt   time.Time `json:"start_at"`
	EndAt     time.Time `json:"end_at"`
	UserId    int       `json:"-"`
}

func GetTaskFilter() *TaskFilter {
	return &TaskFilter{
		Tag:       "",
		Completed: CompletedStatus.Ignored,
		Sprints:   []int{},
		StartAt:   time.Time{},
		EndAt:     time.Time{},
	}
}

func (tf TaskFilter) GetTag() string {
	return tf.Tag
}

func (tf *TaskFilter) SetTag(tag string) ITaskFilter {
	tf.Tag = tag
	return tf
}

func (tf TaskFilter) GetCompleted() int {
	return tf.Completed
}

func (tf *TaskFilter) SetCompleted(completed int) ITaskFilter {
	tf.Completed = completed
	return tf
}

func (tf TaskFilter) GetSprints() *[]int {
	return &tf.Sprints
}

func (tf *TaskFilter) SetSprint(sprints *[]int) ITaskFilter {
	tf.Sprints = *sprints
	return tf
}

func (tf TaskFilter) GetStartAt() time.Time {
	return tf.StartAt
}

func (tf *TaskFilter) SetStartAt(startAt time.Time) ITaskFilter {
	tf.StartAt = startAt
	return tf
}

func (tf TaskFilter) GetEndAt() time.Time {
	return tf.EndAt
}

func (tf *TaskFilter) SetEndAt(endAt time.Time) ITaskFilter {
	tf.EndAt = endAt
	return tf
}

func (tf TaskFilter) GetUserId() int {
	return tf.UserId
}

func (tf *TaskFilter) SetUserId(id int) ITaskFilter {
	tf.UserId = id
	return tf
}
