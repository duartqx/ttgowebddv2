package task

import "time"

type ITaskFilter interface {
	GetTag() string
	SetTag(tag string) ITaskFilter

	GetCompleted() int
	SetCompleted(completed int) ITaskFilter

	GetSprints() *[]int
	SetSprint(sprints *[]int) ITaskFilter

	GetStartAt() time.Time
	SetStartAt(startAt time.Time) ITaskFilter

	GetEndAt() time.Time
	SetEndAt(endAt time.Time) ITaskFilter
}
