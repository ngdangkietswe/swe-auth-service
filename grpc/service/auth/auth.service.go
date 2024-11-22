package auth

import (
	"context"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	"github.com/ngdangkietswe/swe-auth-service/utils"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	"log"
)

type authService struct {
	authRepository repository.IAuthRepository
}

func (a authService) RegisterUser(ctx context.Context, req *auth.User) (*common.UpsertResp, error) {
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
		return nil, err
	}

	req.Password = hashPassword
	entUser, err := a.authRepository.UpsertUser(ctx, req)
	if err != nil {
		log.Fatalf("failed to upsert user: %v", err)
		return nil, err
	}

	return &common.UpsertResp{
		Resp: &common.UpsertResp_Data_{
			Data: &common.UpsertResp_Data{Id: entUser.ID.String()},
		},
	}, nil
}

func (a authService) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthGrpcService(authRepository repository.IAuthRepository) IAuthService {
	return &authService{
		authRepository: authRepository,
	}
}
