package repository

import (
	"sync"
	"time"

	a "github.com/duartqx/ttgowebddv2/internal/domains/auth"
)

type Session struct {
	User      a.ISessionUser
	CreatedAt time.Time
}

func (s Session) GetUser() a.ISessionUser {
	return s.User
}

func (s Session) GetCreatedAt() time.Time {
	return s.CreatedAt
}

type sessionStore struct {
	sessions *map[string]a.ISession
	mutex    *sync.Mutex
}

// Memory Based ISessionRepository implementation
// Could be replaced for a Redis based one
type SessionRepository struct {
	store *sessionStore
}

var sessionRepository *SessionRepository

func GetSessionRepository() *SessionRepository {
	if sessionRepository == nil {
		sessionRepository = &SessionRepository{
			store: &sessionStore{
				sessions: &map[string]a.ISession{},
				mutex:    &sync.Mutex{},
			},
		}
	}
	return sessionRepository
}

func (sr SessionRepository) Get(user a.ISessionUser) (a.ISession, error) {

	sr.store.mutex.Lock()
	defer sr.store.mutex.Unlock()

	session, found := (*sr.store.sessions)[user.GetEmail()]
	if !found {
		createdAt := time.Now()
		session = &Session{User: user, CreatedAt: createdAt}
		(*sr.store.sessions)[user.GetEmail()] = session
	}
	return session, nil
}

func (sr SessionRepository) Set(user a.ISessionUser, createdAt time.Time) error {

	sr.store.mutex.Lock()
	defer sr.store.mutex.Unlock()

	(*sr.store.sessions)[user.GetEmail()] = &Session{
		User: user, CreatedAt: createdAt,
	}
	return nil
}

func (sr SessionRepository) Delete(user a.ISessionUser) error {

	sr.store.mutex.Lock()
	defer sr.store.mutex.Unlock()

	delete((*sr.store.sessions), user.GetEmail())
	return nil
}
