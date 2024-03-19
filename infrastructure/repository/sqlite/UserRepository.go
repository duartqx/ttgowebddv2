package sqlite

import (
	"database/sql"
	"errors"

	sqlbuilder "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"

	e "github.com/duartqx/ddgobase/common/errors"
	u "github.com/duartqx/ddgobase/domains/user"
)

type UserRepository struct {
	db *sqlx.DB
}

func GetUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) FindById(user u.IUser) error {
	sb := sqlbuilder.NewSelectBuilder()

	query, args := sb.Select("*").
		From("users").
		Where(sb.Equal("id", user.GetId())).
		Limit(1).
		Build()

	if err := ur.db.Get(&user, query, args...); err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) FindByEmail(user u.IUser) error {
	sb := sqlbuilder.NewSelectBuilder()

	query, args := sb.Select("*").
		From("users").
		Where(sb.Equal("email", user.GetEmail())).
		Limit(1).
		Build()

	if err := ur.db.Get(&user, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.NotFoundError
		}
		return err
	}
	return nil
}

func (ur UserRepository) ExistsByEmail(email string) bool {
	sb := sqlbuilder.NewSelectBuilder()

	subQuery := sb.Select("1").From("users").Where(sb.Equal("email", email))

	query, args := sb.Select(sb.Exists(subQuery)).Build()

	var exists bool
	if err := ur.db.QueryRow(query, args...).Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
		panic(err)
	}
	return exists
}

func (ur UserRepository) Create(user u.IUser) error {

	sb := sqlbuilder.NewInsertBuilder()

	query, args := sb.InsertInto("users").
		Cols("email", "password", "name").
		Values(user.GetEmail(), user.GetPassword(), user.GetName()).
		SQL("RETURNING id").
		Build()

	var id int
	if err := ur.db.QueryRow(query, args...).Scan(&id); err != nil {
		return err
	}

	user.SetId(id)

	return nil
}

func (ur UserRepository) Update(user u.IUser) error {

	sb := sqlbuilder.NewUpdateBuilder()

	assignements := []string{
		sb.Assign("email", user.GetEmail()),
		sb.Assign("password", user.GetPassword()),
		sb.Assign("name", user.GetName()),
	}

	query, args := sb.Update("users").
		Set(assignements...).
		Where(sb.Equal("id", user.GetId())).
		Build()

	_, err := ur.db.Exec(query, args...)

	return err
}

func (ur UserRepository) Delete(user u.IUser) error {

	sb := sqlbuilder.NewDeleteBuilder()

	query, args := sb.DeleteFrom("users").Where(sb.Equal("id", user.GetId())).Build()

	_, err := ur.db.Exec(query, args...)

	return err
}
