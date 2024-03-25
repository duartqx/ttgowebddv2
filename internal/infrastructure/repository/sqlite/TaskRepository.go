package sqlite

import (
	"database/sql"
	"errors"
	"time"

	sqlb "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"

	e "github.com/duartqx/ttgowebddv2/common/errors"
	t "github.com/duartqx/ttgowebddv2/domains/task"
)

type TaskRepository struct {
	db *sqlx.DB
}

func GetTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (tr TaskRepository) getJoinedQueryBuilder() *sqlb.SelectBuilder {
	sb := sqlb.NewSelectBuilder()
	return sb.
		Select(
			"t.id AS id",
			"t.tag AS tag",
			"t.sprint AS sprint",
			"t.description AS description",
			"t.completed AS completed",
			"t.start_at AS start_at",
			"t.end_at AS end_at",
			"t.user_id AS user_id",
			"COALESCE(u.id, 0) AS 'user.id'",
			"COALESCE(u.name, '') AS 'user.name'",
			"COALESCE(u.email, '') AS 'user.email'",
		).
		From("tasks t").
		Join("users u", "u.id = t.user_id")
}

func (tr TaskRepository) Filter(tf t.ITaskFilter) (*[]t.Task, error) {

	sb := tr.getJoinedQueryBuilder()

	sb.Where(sb.Equal("t.user_id", tf.GetUserId()))

	if tf.GetTag() != "" {
		sb.Where(sb.Equal("t.tag", tf.GetTag()))
	}

	if tf.GetCompleted() != t.CompletedStatus.Ignored {
		sb.Where(sb.Equal("t.completed", tf.GetCompleted()))
	}

	if len(*tf.GetSprints()) != 0 {
		var sprints []interface{}
		for _, sprint := range *tf.GetSprints() {
			sprints = append(sprints, sprint)
		}
		sb.Where(sb.In("t.sprint", sprints...))
	}

	if !tf.GetStartAt().IsZero() {
		sb.Where(sb.Between("t.start_at", tf.GetStartAt(), time.Now()))
	}

	if !tf.GetEndAt().IsZero() {
		sb.Where(sb.Between("t.end_at", tf.GetEndAt(), time.Now()))
	}

	query, args := sb.Build()

	var tasks []t.Task
	if err := tr.db.Select(&tasks, query, args...); err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (tr TaskRepository) Update(task *t.Task) error {
	ub := sqlb.NewUpdateBuilder()

	assignments := []string{
		ub.Assign("tag", task.GetTag()),
		ub.Assign("sprint", task.GetSprint()),
		ub.Assign("description", task.GetDescription()),
		ub.Assign("completed", task.GetCompleted()),
	}

	if !task.GetStartAt().IsZero() {
		assignments = append(assignments, ub.Assign("start_at", task.GetStartAt()))
	}

	if !task.GetEndAt().IsZero() {
		assignments = append(assignments, ub.Assign("end_at", task.GetEndAt()))
	}

	query, args := ub.Update("tasks").
		Set(assignments...).
		Where(ub.Equal("id", task.GetId())).
		Build()

	if _, err := tr.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}

func (tr TaskRepository) findByWhere(task *t.Task, query string, args ...interface{}) error {

	if err := tr.db.Get(&task, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.NotFoundError
		}
		return err
	}

	return nil
}

func (tr TaskRepository) FindById(task *t.Task) error {
	sb := tr.getJoinedQueryBuilder()
	query, args := sb.Where(sb.Equal("t.id", task.GetId())).Build()
	return tr.findByWhere(task, query, args...)
}

func (tr TaskRepository) FindByTag(task *t.Task) error {
	sb := tr.getJoinedQueryBuilder()
	query, args := sb.Where(sb.Equal("t.tag", task.GetTag())).Build()
	return tr.findByWhere(task, query, args...)
}

func (tr TaskRepository) Create(task *t.Task) error {

	args := []interface{}{
		task.GetTag(),
		task.GetSprint(),
		task.GetDescription(),
		task.GetUserId(),
	}

	query := `
		BEGIN TRANSACTION;

		WITH new_task AS (
			INSERT INTO tasks (tag, sprint, description, user_id)
			VALUES ($1, $2, $3, $4)
			RETURNING id, start_at, user_id
		);

		SELECT
			t.id AS id,
			t.start_at AS start_at,
			u.id AS 'user.id',
			u.name AS 'user.name',
			u.email AS 'user.email'
		FROM new_task t
		LEFT JOIN users u ON u.id = t.user_id;

		COMMIT;
	`

	if err := tr.db.Get(task, query, args...); err != nil {
		return err
	}

	return nil
}

func (tr TaskRepository) GetSprints() *[]string {

	var sprints []string

	query, _ := sqlb.NewSelectBuilder().From("tasks").Select("sprint").GroupBy("sprint").Build()

	if err := tr.db.Select(&sprints, query); err != nil {
		panic(err)
	}

	return &sprints
}
