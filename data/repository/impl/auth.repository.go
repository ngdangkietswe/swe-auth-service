package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/ent/user"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	"github.com/ngdangkietswe/swe-auth-service/utils"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type authRepository struct {
	entClient *ent.Client
}

// EnableOrDisable2FA is a function that enables or disables 2FA for a user
func (a authRepository) EnableOrDisable2FA(ctx context.Context, userId string, enable bool) (*ent.User, error) {
	query := a.entClient.User.UpdateOneID(uuid.MustParse(userId)).SetEnable2fa(enable)

	if enable {
		query.SetSecret2fa(utils.GenerateSecret())
	} else {
		query.SetSecret2fa("")
	}

	entUser, err := query.Save(ctx)

	return entUser, err
}

// FindByUsernameOrEmail is a function that finds a user by username or email
func (a authRepository) FindByUsernameOrEmail(ctx context.Context, username string, email string) (*ent.User, error) {
	entUser, err := a.entClient.User.Query().Where(user.Or(user.UsernameEQ(username), user.EmailEQ(email))).First(ctx)
	return entUser, err
}

// FindByUsername is a function that finds a user by username
func (a authRepository) FindByUsername(ctx context.Context, username string) (*ent.User, error) {
	entUser, err := a.entClient.User.Query().Where(user.UsernameEQ(username)).First(ctx)
	return entUser, err
}

// UpsertUser is a function that upserts a user. If the user has an ID, it will update the user. Otherwise, it will create a new user
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
