package task

type ITaskRepository interface {
	Filter(tf ITaskFilter) (*[]Task, error)
	Create(task Task) error
	Update(task Task) error

	FindById(task Task) error
	FindByTag(task Task) error

	GetSprints() *[]string
}
