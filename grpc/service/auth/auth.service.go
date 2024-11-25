package auth

import (
	"context"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/auth"
	"github.com/ngdangkietswe/swe-auth-service/kafka/constant"
	"github.com/ngdangkietswe/swe-auth-service/kafka/model"
	"github.com/ngdangkietswe/swe-auth-service/kafka/producer"
	"github.com/ngdangkietswe/swe-auth-service/utils"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	"time"
)

type authService struct {
	authRepository repository.IAuthRepository
	authValidator  validator.IAuthValidator
}

func (a authService) RegisterUser(ctx context.Context, req *auth.User) (*common.UpsertResp, error) {
	// Validate request
	err := a.authValidator.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = hashPassword
	entUser, err := a.authRepository.UpsertUser(ctx, req)
	if err != nil {
		return nil, err
	}

	// Produce message to kafka. This message will be consumed by swe notification service
	kProducer := producer.NewKProducer(constant.TopicRegisterUser)
	go func() {
		registerUser := model.RegisterUser{
			Username:  req.Username,
			Email:     req.Email,
			CreatedAt: time.Now().Format(time.RFC3339),
		}
		kProducer.Produce(entUser.ID.String(), registerUser)
	}()

	return &common.UpsertResp{
		Success: true,
		Resp: &common.UpsertResp_Data_{
			Data: &common.UpsertResp_Data{Id: entUser.ID.String()},
		},
	}, nil
}

func (a authService) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	var (
		entUser *ent.User
		err     error
		token   string
	)

	// Check if user exists with username
	entUser, err = a.authRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return &auth.LoginResp{
			Error: &common.Error{
				Code:    401,
				Message: "Username or password is incorrect",
			},
		}, nil
	}

	// Compare password
	if err := utils.CheckPasswordHash(entUser.Password, req.Password); err != nil {
		return &auth.LoginResp{
			Error: &common.Error{
				Code:    401,
				Message: "Username or password is incorrect",
			},
		}, nil
	}

	// Generate token
	token, err = utils.GenerateToken(entUser)
	if err != nil {
		return &auth.LoginResp{
			Error: &common.Error{
				Code:    401,
				Message: "Unknown error",
			},
		}, nil
	}

	return &auth.LoginResp{
		Token: token,
	}, nil
}

func NewAuthGrpcService(
	authRepository repository.IAuthRepository,
	authValidator validator.IAuthValidator) IAuthService {
	return &authService{
		authRepository: authRepository,
		authValidator:  authValidator,
	}
}
