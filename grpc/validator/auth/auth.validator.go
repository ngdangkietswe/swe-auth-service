package auth

import (
	"context"
	"errors"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type authValidator struct {
	authRepository repository.IAuthRepository
}

// RegisterUser is a function that validates the user registration request
func (a authValidator) RegisterUser(ctx context.Context, req *auth.User) error {
	if req.Username == "" {
		return errors.New("username is required")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}

	entUser, _ := a.authRepository.FindByUsernameOrEmail(ctx, req.Username, req.Email)
	if entUser != nil {
		return errors.New("username or email is existed")
	}

	return nil
}

func NewAuthValidator(authRepository repository.IAuthRepository) IAuthValidator {
	return authValidator{
		authRepository: authRepository,
	}
}
