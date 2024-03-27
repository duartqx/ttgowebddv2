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
			name:    "EmptyFilterHasFiveResults",
			filter:  task.GetTaskFilter(),
			results: 5,
		},
		{
			name:    "CompletedFilterHasTwoResults",
			filter:  task.GetTaskFilter().SetCompleted(task.CompletedStatus.Completed),
			results: 2,
		},
		{
			name:    "Sprint81FilterHasThreeResults",
			filter:  task.GetTaskFilter().SetSprint(&[]int{81}),
			results: 3,
		},
		{
			name:    "Sprint81And82FilterHasFourResults",
			filter:  task.GetTaskFilter().SetSprint(&[]int{81, 82}),
			results: 4,
		},
		{
			name:    "UserId3FilterHasOneResults",
			filter:  task.GetTaskFilter().SetUserId(3),
			results: 1,
		},
		{
			name:    "UserId2AndCompletedFilterHasOneResults",
			filter:  task.GetTaskFilter().SetUserId(3).SetCompleted(task.CompletedStatus.Completed),
			results: 1,
		},
		{
			name:    "TagAJABCDFilterHasOneResults",
			filter:  task.GetTaskFilter().SetTag("AJ-ABCD"),
			results: 1,
		},
		{
			name:    "Sprint84FilterHasZeroResults",
			filter:  task.GetTaskFilter().SetSprint(&[]int{84}),
			results: 0,
		},
		{
			name:    "UserId4FilterHasZeroResults",
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
				t.Fatalf("Mismatch results, have %d, wants %d!", results, tt.results)
			}
		})
	}
}
