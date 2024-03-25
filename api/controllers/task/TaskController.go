package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	dto "github.com/duartqx/ttgowebddv2/api/dto"
	h "github.com/duartqx/ttgowebddv2/api/http"
	e "github.com/duartqx/ttgowebddv2/common/errors"
	a "github.com/duartqx/ttgowebddv2/domains/auth"
	t "github.com/duartqx/ttgowebddv2/domains/task"
)

type TaskController struct {
	taskService    t.ITaskService
	sessionService a.ISessionService
}

var taskController *TaskController

func GetTaskController(
	taskService t.ITaskService, sessionService a.ISessionService,
) *TaskController {
	if taskController == nil {
		taskController = &TaskController{
			taskService:    taskService,
			sessionService: sessionService,
		}
	}
	return taskController
}

func (tc TaskController) Create(w http.ResponseWriter, r *http.Request) {

	user := tc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	var taskDTO dto.TaskCreateDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	task := t.GetNewTask()

	task.
		SetUser(user.ToUser()).
		SetTag(taskDTO.Tag).
		SetSprint(fmt.Sprintf("%d", taskDTO.Sprint)).
		SetDescription(taskDTO.Description)

	if err := tc.taskService.Create(task); err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusCreated, task)
}

func (tc TaskController) Update(w http.ResponseWriter, r *http.Request) {
	user := tc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	taskId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}

	var taskDTO dto.TaskUpdateDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	task := t.GetNewTask()

	task.
		SetId(taskId).
		SetUser(user.ToUser()).
		SetTag(taskDTO.Tag).
		SetSprint(fmt.Sprintf("%d", taskDTO.Sprint)).
		SetDescription(taskDTO.Description).
		SetStartAt(&taskDTO.StartAt).
		SetEndAt(&taskDTO.EndAt).
		SetCompleted(taskDTO.Completed)

	if err := tc.taskService.Create(task); err != nil {
		h.ErrorResponse(w, err)
		return
	}

	h.JsonResponse(w, http.StatusOK, task)
}

func (tc TaskController) Filter(w http.ResponseWriter, r *http.Request) {

	user := tc.sessionService.GetSessionUser(r.Context())
	if user == nil {
		http.Error(w, e.ForbiddenError.Error(), http.StatusForbidden)
		return
	}

	taskFilter := t.GetTaskFilter()

	if err := json.NewDecoder(r.Body).Decode(&taskFilter); err != nil {
		http.Error(w, e.BadRequestError.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	taskFilter.SetUserId(user.GetId())

	tasks, err := tc.taskService.Filter(taskFilter)
	if err != nil {
		panic(err)
	}

	h.JsonResponse(w, http.StatusOK, tasks)
}
