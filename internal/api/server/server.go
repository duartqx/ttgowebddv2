package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"

	ac "github.com/duartqx/ttgowebddv2/internal/api/controllers/auth"
	tc "github.com/duartqx/ttgowebddv2/internal/api/controllers/task"
	uc "github.com/duartqx/ttgowebddv2/internal/api/controllers/user"
	as "github.com/duartqx/ttgowebddv2/internal/application/services/auth"
	ts "github.com/duartqx/ttgowebddv2/internal/application/services/task"
	us "github.com/duartqx/ttgowebddv2/internal/application/services/user"
	a "github.com/duartqx/ttgowebddv2/internal/domains/auth"
	t "github.com/duartqx/ttgowebddv2/internal/domains/task"
	u "github.com/duartqx/ttgowebddv2/internal/domains/user"

	cm "github.com/duartqx/ddgomiddlewares/cors"
	lm "github.com/duartqx/ddgomiddlewares/logger"
	rm "github.com/duartqx/ddgomiddlewares/recovery"
)

type ServerConfig struct {
	Db        *sqlx.DB
	JwtSecret []byte
	Cors      bool

	UserRepository    u.IUserRepository
	SessionRepository a.ISessionRepository
	TaskRepository    t.ITaskRepository
}

type server struct {
	db        *sqlx.DB
	jwtSecret []byte
	cors      bool

	mux *http.ServeMux

	userRepository    u.IUserRepository
	sessionRepository a.ISessionRepository
	taskRepository    t.ITaskRepository

	jwtController *ac.JwtController
}

func NewServer(config *ServerConfig) *server {

	jwtService := as.GetJwtAuthService(
		config.UserRepository, config.SessionRepository, &config.JwtSecret,
	)
	jwtController := ac.GetJwtController(jwtService)

	return &server{
		db:        config.Db,
		jwtSecret: config.JwtSecret,
		cors:      config.Cors,

		mux: http.NewServeMux(),

		userRepository:    config.UserRepository,
		sessionRepository: config.SessionRepository,
		taskRepository:    config.TaskRepository,

		jwtController: jwtController,
	}
}

func (s *server) AddBaseUserRoutes() *server {

	userController := uc.GetUserController(
		us.GetUserService(s.userRepository),
		as.GetSessionService(s.sessionRepository),
	)

	userMux := http.NewServeMux()

	userMux.Handle(
		"POST /login/{$}",
		s.jwtController.UnauthenticatedMiddleware(
			http.HandlerFunc(s.jwtController.Login),
		),
	)

	userMux.Handle(
		"DELETE /logout/{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(s.jwtController.Logout),
		),
	)

	userMux.Handle(
		"GET /{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(userController.Get),
		),
	)

	userMux.Handle(
		"POST /{$}",
		s.jwtController.UnauthenticatedMiddleware(
			http.HandlerFunc(userController.Create),
		),
	)

	userMux.Handle(
		"PATCH /{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(userController.UpdatePassword),
		),
	)

	s.AddGroup("/api/users/", userMux)

	return s
}

func (s *server) AddTaskRoutes() *server {
	taskController := tc.GetTaskController(
		ts.GetTaskService(s.taskRepository),
		as.GetSessionService(s.sessionRepository),
	)

	taskMux := http.NewServeMux()

	taskMux.Handle(
		"POST /{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(taskController.Create),
		),
	)

	taskMux.Handle(
		"PATCH /{id}/{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(taskController.Update),
		),
	)

	taskMux.Handle(
		"POST /filter/{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(taskController.Filter),
		),
	)

	taskMux.Handle(
		"GET /sprints/{$}",
		s.jwtController.AuthenticatedMiddleware(
			http.HandlerFunc(taskController.Sprints),
		),
	)

	s.AddGroup("/api/tasks/", taskMux)

	return s
}

func (s server) Use(
	mux http.Handler, middlewares ...func(http.Handler) http.Handler,
) http.Handler {
	wrapped := mux
	for _, middleware := range middlewares {
		wrapped = middleware(wrapped)
	}
	return wrapped
}

func (s *server) AddGroup(pattern string, handler http.Handler) error {
	if !strings.HasPrefix(pattern, "/") && !strings.HasSuffix(pattern, "/") {
		return fmt.Errorf("Invalid Pattern")
	}

	prefix := strings.TrimSuffix(pattern, "/")

	s.mux.Handle(pattern, http.StripPrefix(prefix, handler))

	return nil
}

func (s *server) Build() http.Handler {

	s.AddBaseUserRoutes().AddTaskRoutes()

	mux := http.Handler(s.mux)

	if s.cors {
		mux = s.Use(mux, cm.CorsMiddleware)
	}

	return s.Use(
		mux,
		lm.LoggerMiddleware,
		rm.RecoveryMiddleware,
	)
}
