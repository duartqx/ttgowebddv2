package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"

	"github.com/duartqx/ttgowebddv2/internal/api/server"
	"github.com/duartqx/ttgowebddv2/internal/infrastructure/repository"
)

//go:embed internal/presentation/react/dist/*
var assets embed.FS

type env struct {
	dbStr  string `validate:"required"`
	secret string `validate:"required"`
	addr   string `validate:"required"`
	debug  bool
}

func GetEnv() *env {
	e := &env{
		dbStr:  os.Getenv("CONNECTION_STR"),
		secret: os.Getenv("SECRET"),
		addr:   os.Getenv("SERVER_ADDR"),
		debug:  os.Getenv("DEBUG") == "1",
	}
	if errs := validator.New().Struct(e); errs != nil {
		log.Fatalln(errs)
	}
	return e
}

func GetServer(db *sqlx.DB, env *env) http.Handler {
	return server.
		NewServer(&server.ServerConfig{
			Db:                db,
			JwtSecret:         []byte(env.secret),
			Cors:              env.debug,
			UserRepository:    getUserRepository(db),
			SessionRepository: repository.GetSessionRepository(),
			TaskRepository:    getTaskRepository(db),
			AssetsFS:          assets,
		}).
		Build()
}

func main() {

	env := GetEnv()

	db := initDB(env.dbStr) // Defined at compile time
	defer db.Close()

	mux := GetServer(db, env)

	srv := &http.Server{
		Handler:      mux,
		Addr:         env.addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Listening and Serving at:", env.addr)

	c := make(chan os.Signal, 1)
	// Graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(c, os.Interrupt)

	// Block until we interrupt signal
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)

	os.Exit(0)
}
