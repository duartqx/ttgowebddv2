package task_test

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"

	ts "github.com/duartqx/ttgowebddv2/src/application/services/task"
	"github.com/duartqx/ttgowebddv2/src/domains/task"
	r "github.com/duartqx/ttgowebddv2/src/infrastructure/repository/sqlite"
)

var (
	db *sqlx.DB

	taskRepository task.ITaskRepository
	taskService    task.ITaskService
)

func TestMain(m *testing.M) {

	db = r.GetInMemoryDB("taskservice")
	defer db.Close()

	taskRepository = r.GetTaskRepository(db)
	taskService = ts.GetTaskService(taskRepository)

	r.Seed(db)

	code := m.Run()

	os.Exit(code)
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name    string
		filter  task.ITaskFilter
		results int
	}{
		{
			name:    "EmptyFilterHasFiveResultsTestCase",
			filter:  task.GetTaskFilter(),
			results: 5,
		},
		{
			name: "CompletedFilterHasTwoResultsTestCase",
			filter: task.GetTaskFilter().
				SetCompleted(task.CompletedStatus.Completed),
			results: 2,
		},
		{
			name:    "Sprint81FilterHasThreeResultsTestCase",
			filter:  task.GetTaskFilter().SetSprint(&[]int{81}),
			results: 3,
		},
		{
			name:    "Sprint81And82FilterHasFourResultsTestCase",
			filter:  task.GetTaskFilter().SetSprint(&[]int{81, 82}),
			results: 4,
		},
		{
			name:    "UserId3FilterHasOneResultsTestCase",
			filter:  task.GetTaskFilter().SetUserId(3),
			results: 1,
		},
		{
			name: "UserId2AndCompletedFilterHasOneResultsTestCase",
			filter: task.GetTaskFilter().
				SetUserId(3).
				SetCompleted(task.CompletedStatus.Completed),
			results: 1,
		},
		{
			name:    "TagAJABCDFilterHasOneResultsTestCase",
			filter:  task.GetTaskFilter().SetTag("AJ-ABCD"),
			results: 1,
		},
		{
			name:    "Sprint84FilterHasZeroResultsTestCase",
			filter:  task.GetTaskFilter().SetSprint(&[]int{84}),
			results: 0,
		},
		{
			name:    "UserId4FilterHasZeroResultsTestCase",
			filter:  task.GetTaskFilter().SetUserId(4),
			results: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks, err := taskService.Filter(tt.filter)

			if err != nil {
				t.Fatal(err)
			}

			results := len(*tasks)
			if results != tt.results {
				t.Fatalf(
					"Mismatch results, have %d, wants %d!",
					results,
					tt.results,
				)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name string
		task *task.Task
		err  bool
	}{
		{
			name: "InvalidTagAndSprintAndUserIdTestCase",
			task: task.GetNewTask(),
			err:  true,
		},
		{
			name: "InvalidTagAndUserIdTestCase",
			task: task.GetNewTask().SetSprint("81"),
			err:  true,
		},
		{
			name: "InvalidTagTestCase",
			task: task.GetNewTask().SetSprint("81").SetUserId(1),
			err:  true,
		},
		{
			name: "InvalidUserIdTestCase",
			task: task.GetNewTask().SetSprint("81").SetTag("AJ-ABCD"),
			err:  true,
		},
		{
			name: "UserIdDoesNotExistsTestCase",
			task: task.GetNewTask().SetSprint("81").SetTag("AJ-ABCD").SetUserId(19),
			err:  true,
		},
		{
			name: "ValidTestCase1",
			task: task.GetNewTask().SetSprint("81").SetTag("AJ-EFGH").SetUserId(1),
			err:  false,
		},
		{
			name: "ValidTestCase2",
			task: task.GetNewTask().SetSprint("82").SetTag("AJ-IJKL").SetUserId(2),
			err:  false,
		},
		{
			name: "ValidTestCase3",
			task: task.GetNewTask().SetSprint("83").SetTag("AJ-MNOP").SetUserId(3),
			err:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := taskService.Create(tt.task)

			t.Logf("Error: %v", err)

			if tt.err && err == nil {
				t.Fatal("Should've failed but error was Nil")
			} else if !tt.err && err != nil {
				t.Fatalf("Shouldn't have failed but got %s", err.Error())
			}

			if !tt.err && tt.task.GetId() == 0 {
				t.Fatalf("The Task.Id did not updated after Create")
			} else if tt.task.GetId() != 0 {
				t.Logf(
					"Created Task Id: %d, Start At: %v",
					tt.task.GetId(),
					tt.task.GetStartAt(),
				)
			}
		})
	}

}
