package auth

import (
	"context"
	"github.com/google/wire"
	"github.com/ngdangkietswe/swe-auth-service/data/datasource"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type IAuthRepository interface {
	ChangePassword(ctx context.Context, id, newPassword string) (*ent.User, error)
	UpsertUser(ctx context.Context, user *auth.User) (*ent.User, error)
	FindById(ctx context.Context, id string) (*ent.User, error)
	FindByUsername(ctx context.Context, username string) (*ent.User, error)
	FindByUsernameOrEmail(ctx context.Context, username string, email string) (*ent.User, error)
	EnableOrDisable2FA(ctx context.Context, userId string, enable bool) (*ent.User, error)
	ExistsById(ctx context.Context, id string) (bool, error)
}

// ProvideAuthRepository is a function to provide an auth repository
func ProvideAuthRepository() (a IAuthRepository) {
	wire.Build(
		datasource.ProvideEntClient,
		NewAuthRepository,
	)
	return
}
