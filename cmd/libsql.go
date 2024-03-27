//go:build libsql
// +build libsql

package main

import (
	"log"

	"github.com/duartqx/ddgobase/infrastructure/src/repository/sqlite"
	"github.com/jmoiron/sqlx"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func initDB(conn string) *sqlx.DB {
	db, err := sqlx.Connect("libsql", conn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func getUserRepository(db *sqlx.DB) *sqlite.UserRepository {
	return sqlite.GetUserRepository(db)
}
