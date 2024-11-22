package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/user"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type authRepository struct {
	entClient *ent.Client
}

func (a authRepository) FindByUsername(ctx context.Context, username string) (*ent.User, error) {
	entUser, err := a.entClient.User.Query().Where(user.UsernameEQ(username)).First(ctx)
	return entUser, err
}

func (a authRepository) UpsertUser(ctx context.Context, user *auth.User) (*ent.User, error) {
	var entUser *ent.User
	var err error

	if user.Id != nil {
		entUser, err = a.entClient.User.UpdateOneID(uuid.MustParse(*user.Id)).
			SetUsername(user.Username).
			SetEmail(user.Email).
			SetPassword(user.Password).
			Save(ctx)
	} else {
		entUser, err = a.entClient.User.Create().
			SetUsername(user.Username).
			SetEmail(user.Email).
			SetPassword(user.Password).
			Save(ctx)
	}

	return entUser, err
}

func NewAuthRepository(entClient *ent.Client) repository.IAuthRepository {
	return &authRepository{
		entClient: entClient,
	}
}