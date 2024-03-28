package sqlite

import (
	"database/sql"
	"errors"
	"log"

	sqlbuilder "github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"

	e "github.com/duartqx/ttgowebddv2/internal/common/errors"
	u "github.com/duartqx/ttgowebddv2/internal/domains/user"
)

type UserRepository struct {
	db *sqlx.DB
}

func GetUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) FindById(user *u.User) error {
	sb := sqlbuilder.NewSelectBuilder()

	query, args := sb.Select("*").
		From("users").
		Where(sb.Equal("id", user.GetId())).
		Limit(1).
		Build()

	if err := ur.db.Get(user, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.NotFoundError
		}
		return err
	}
	return nil
}

func (ur UserRepository) FindByEmail(user *u.User) error {
	sb := sqlbuilder.NewSelectBuilder()

	query, args := sb.Select("*").
		From("users").
		Where(sb.Equal("email", user.GetEmail())).
		Limit(1).
		Build()

	if err := ur.db.Get(user, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.NotFoundError
		}
		return err
	}

	return nil
}

func (ur UserRepository) ExistsByEmail(email string) bool {

	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)"

	var exists bool
	if err := ur.db.QueryRow(query, email).Scan(&exists); err != nil {
		log.Println(err)
	}
	return exists
}

func (ur UserRepository) Create(user *u.User) (err error) {

	sb := sqlbuilder.NewInsertBuilder()

	query, args := sb.InsertInto("users").
		Cols("email", "password", "name").
		Values(user.GetEmail(), user.GetPassword(), user.GetName()).
		SQL("RETURNING id").
		Build()

	if err := ur.db.QueryRow(query, args...).Scan(&user.Id); err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) Update(user *u.User) (err error) {

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

	if _, err := ur.db.Exec(query, args...); err != nil {
		return err
	}

	return err
}

func (ur UserRepository) Delete(user *u.User) (err error) {

	sb := sqlbuilder.NewDeleteBuilder()

	query, args := sb.DeleteFrom("users").Where(sb.Equal("id", user.GetId())).Build()

	if _, err := ur.db.Exec(query, args...); err != nil {
		return err
	}

	return err
}
