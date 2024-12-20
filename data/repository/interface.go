package repository

import (
	"context"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type (
	IAuthRepository interface {
		UpsertUser(ctx context.Context, user *auth.User) (*ent.User, error)
		FindByUsername(ctx context.Context, username string) (*ent.User, error)
		FindByUsernameOrEmail(ctx context.Context, username string, email string) (*ent.User, error)
		EnableOrDisable2FA(ctx context.Context, userId string, enable bool) (*ent.User, error)
	}
)
