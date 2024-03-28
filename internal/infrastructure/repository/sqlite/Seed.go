package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func GetInMemoryDB(name string) *sqlx.DB {
	db, err := sqlx.Open(
		"sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared", name),
	)
	if err != nil {
		panic(err)
	}
	return db
}

func Seed(db *sqlx.DB) error {
	_, err := db.Exec(
		`
			BEGIN TRANSACTION;
			CREATE TABLE users (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL UNIQUE,
				password VARCHAR(255) NOT NULL
			);

			INSERT INTO users ( name, email, password )
			VALUES
				( 'Test User 1', 'test1@test1.com', '$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW' ),
				( 'Test User 2', 'test2@test2.com', '$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW' ),
				( 'Test User 3', 'test3@test3.com', '$2a$10$HpNsS.a6Q6ThR0nsAuuMS.6UbSGDB9/Do5C.zZFfJBEKjOQOk/UaW' );

			CREATE TABLE tasks (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				tag TEXT NOT NULL,
				description TEXT NOT NULL,
				start_at DATETIME DEFAULT (datetime('now')),
				end_at DATETIME DEFAULT NULL,
				completed INTEGER DEFAULT 0,
				sprint TEXT NOT NULL,
				user_id INTEGER NOT NULL DEFAULT 1 REFERENCES users(id) ON DELETE CASCADE ON UPDATE NO ACTION
			);
			CREATE INDEX tasks_sprint_index ON tasks (sprint ASC);

			COMMIT;
		`,
	)

	return err
}
