package task

import (
	"fmt"

	e "github.com/duartqx/ttgowebddv2/src/common/errors"
	t "github.com/duartqx/ttgowebddv2/src/domains/task"
)

type TaskService struct {
	taskRepository t.ITaskRepository
}

var taskService *TaskService

func GetTaskService(taskRepository t.ITaskRepository) *TaskService {
	if taskService == nil {
		taskService = &TaskService{taskRepository: taskRepository}
	}
	return taskService
}

func (ts TaskService) Filter(tf t.ITaskFilter) (*[]t.Task, error) {
	tasks, err := ts.taskRepository.Filter(tf)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts TaskService) isInvalidTask(task *t.Task) bool {
	return task.GetTag() == "" || task.GetSprint() == "" || task.GetUserId() == 0
}

func (ts TaskService) Create(task *t.Task) error {

	if ts.isInvalidTask(task) {
		return fmt.Errorf("%w: Invalid Task", e.BadRequestError)
	}

	if err := ts.taskRepository.Create(task); err != nil {
		return err
	}

	return nil
}

func (ts TaskService) Update(task *t.Task) error {
	if task.GetId() == 0 || ts.isInvalidTask(task) {
		return fmt.Errorf("%w: Invalid Task", e.BadRequestError)
	}

	if err := ts.taskRepository.Update(task); err != nil {
		return fmt.Errorf("%w: Error Trying to Update Task", e.InternalError)
	}

	return nil
}

func (ts TaskService) GetSprints() *[]string {
	return ts.taskRepository.GetSprints()
}
