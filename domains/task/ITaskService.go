package task

type ITaskService interface {
	Filter(tf ITaskFilter) (*[]ITask, error)
	Create(task ITask) error
	Update(task ITask) error

	GetSprints() *[]int
}
