package auth

import (
	"context"

	a "github.com/duartqx/ddgobase/internal/domains/auth"
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
	contextUser := ctx.Value("user")

	if contextUser == nil {
		return nil
	}

	user, ok := contextUser.(*a.SessionUser)
	if !ok {
		return nil
	}

	session, _ := ss.sessionRepository.Get(user)
	return session.GetUser()
}
