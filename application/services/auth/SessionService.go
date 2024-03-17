package auth

import (
	"context"

	a "github.com/duartqx/ddgobase/domains/auth"
)

type SessionService struct {
	sessionRepository a.ISessionRepository
}

var sessionService *SessionService

func GetSessionService(sessionRepository a.ISessionRepository) *SessionService {
	if sessionService == nil {
		sessionService = &SessionService{
			sessionRepository: sessionRepository,
		}
	}
	return sessionService
}

func (ss SessionService) GetSessionUser(ctx context.Context) a.ISessionUser {
	user, ok := ctx.Value("user").(*a.SessionUser)
	if !ok {
		return nil
	}
	session, _ := ss.sessionRepository.Get(user)
	return session.GetUser()
}
