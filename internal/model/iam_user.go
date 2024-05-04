package model

import (
	"context"

	"github.com/binus-thesis-team/iam-service/auth"
)

type SessionUser struct {
	*auth.User
}

// GetUserID get userID
func (s *SessionUser) GetUserID() int64 {
	if s.User == nil {
		return 0
	}
	return s.User.ID
}

// GetUserFromCtx get auth user from context
func GetUserFromCtx(ctx context.Context) SessionUser {
	return SessionUser{
		User: auth.GetUserFromCtx(ctx),
	}
}
