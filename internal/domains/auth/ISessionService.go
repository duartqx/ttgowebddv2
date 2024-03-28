package auth

import "context"

type ISessionService interface {
	GetSessionUser(ctx context.Context) ISessionUser
}
