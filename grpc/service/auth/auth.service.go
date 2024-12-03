package auth

import (
	"context"
	"errors"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	"github.com/ngdangkietswe/swe-auth-service/grpc/mapper"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/auth"
	"github.com/ngdangkietswe/swe-auth-service/kafka/constant"
	"github.com/ngdangkietswe/swe-auth-service/kafka/model"
	"github.com/ngdangkietswe/swe-auth-service/kafka/producer"
	"github.com/ngdangkietswe/swe-auth-service/utils"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/grpc/domain"
	grpcutil "github.com/ngdangkietswe/swe-go-common-shared/grpc/util"
	commonutil "github.com/ngdangkietswe/swe-go-common-shared/util"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	"time"
)

type authService struct {
	authRepository repository.IAuthRepository
	authValidator  validator.IAuthValidator
}

// EnableOrDisable2FA is a function that enables or disables 2FA for a user.
func (a authService) EnableOrDisable2FA(ctx context.Context, req *auth.EnableOrDisable2FAReq) (*auth.EnableOrDisable2FAResp, error) {
	principal := grpcutil.GetGrpcPrincipal(ctx)
	userId := principal.UserId
	entUser, err := a.authRepository.EnableOrDisable2FA(ctx, userId, req.Enable)
	if err != nil {
		return nil, errors.New("user not found")
	}

	resp := &auth.EnableOrDisable2FAResp{
		Success: true,
	}

	if req.Enable {
		resp.QrCodeImageUrl = utils.GenerateTOTPWithSecret(*entUser.Secret2fa)
	}

	return resp, nil
}

// RegisterUser is a function that registers a new user.
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

// Login is a function that logs in a user.
func (a authService) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	var (
		entUser *ent.User
		err     error
		token   string
	)

	// Check if user exists with username
	entUser, err = a.authRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return mapper.AsFailed("Username or password is incorrect")
	}

	// Compare password
	if err := utils.CheckPasswordHash(entUser.Password, req.Password); err != nil {
		return mapper.AsFailed("Username or password is incorrect")
	}

	// Validate 2fa if it is enabled for the user
	if entUser.Enable2fa {
		if req.Otp == nil || *req.Otp == "" {
			return mapper.AsFailed("Two-factor authentication is required")
		}
		if !utils.VerifyOTP(*entUser.Secret2fa, *req.Otp) {
			return mapper.AsFailed("Two-factor authentication is incorrect")
		}
	}

	grpcUser := &domain.GrpcUser{
		Id:       entUser.ID.String(),
		Username: entUser.Username,
		Email:    entUser.Email,
	}

	// Generate token
	token, err = commonutil.GenerateToken(grpcUser, false)
	if err != nil {
		return mapper.AsFailed("Unknown error")
	}

	// Generate refresh token
	refreshToken, err := commonutil.GenerateToken(grpcUser, true)
	if err != nil {
		return mapper.AsFailed("Unknown error")
	}

	return &auth.LoginResp{
		Success: true,
		Resp: &auth.LoginResp_Data_{
			Data: &auth.LoginResp_Data{
				AccessToken:           token,
				AccessTokenExpiresIn:  config.GetString("JWT_EXPIRATION", ""),
				RefreshToken:          refreshToken,
				RefreshTokenExpiresIn: config.GetString("REFRESH_TOKEN_EXPIRATION", ""),
				TokenType:             "Bearer",
				TwoFactorAuth:         entUser.Enable2fa,
			},
		},
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
