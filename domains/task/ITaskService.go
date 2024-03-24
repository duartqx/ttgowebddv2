package task

type ITaskService interface {
	Filter(tf ITaskFilter) (*[]Task, error)
	Create(task *Task) error
	Update(task *Task) error

	GetSprints() *[]string
}
