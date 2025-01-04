package grpcserviceauth

import (
	"context"
	"errors"
	"fmt"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	authrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/auth"
	"github.com/ngdangkietswe/swe-auth-service/grpc/mapper"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/auth"
	"github.com/ngdangkietswe/swe-auth-service/kafka/producer"
	"github.com/ngdangkietswe/swe-auth-service/utils"
	"github.com/ngdangkietswe/swe-go-common-shared/cache"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/constants"
	"github.com/ngdangkietswe/swe-go-common-shared/domain"
	grpcdomain "github.com/ngdangkietswe/swe-go-common-shared/grpc/domain"
	grpcutil "github.com/ngdangkietswe/swe-go-common-shared/grpc/util"
	commonutil "github.com/ngdangkietswe/swe-go-common-shared/util"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	"time"
)

type authService struct {
	redisCache     *cache.RedisCache
	authRepository authrepo.IAuthRepository
	authValidator  validator.IAuthValidator
}

// ForgotPassword is a function that sends an email to a user to reset the password.
func (a authService) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordReq) (*common.EmptyResp, error) {
	exists, err := a.authRepository.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("user not found")
	}

	token, err := commonutil.GenerateSecureToken()
	if err != nil {
		return nil, err
	}

	resetPassword := domain.ResetPassword{
		Email: req.Email,
		Token: token,
	}

	// Cache token in redis with expiration time in 30 minutes
	cacheKey := fmt.Sprintf("%s_%s", constants.ResetPasswordCacheKeyPrefix, token)
	err = a.redisCache.Set(cacheKey, resetPassword, time.Minute*30)
	if err != nil {
		return nil, err
	}

	// Produce message to kafka. This message will be consumed by swe notification service
	kProducer := producer.NewKProducer(constants.TopicResetPassword)
	go func() {
		kProducer.Produce(token, resetPassword)
	}()

	return &common.EmptyResp{
		Success: true,
	}, nil
}

// ResetPassword is a function that resets the password of a user.
func (a authService) ResetPassword(ctx context.Context, req *auth.ResetPasswordReq) (*common.EmptyResp, error) {
	var resetPassword *domain.ResetPassword

	// Validate request
	err := a.authValidator.ResetPassword(req)
	if err != nil {
		return nil, err
	}

	// Check if token exists in redis
	cacheKey := fmt.Sprintf("%s_%s", constants.ResetPasswordCacheKeyPrefix, req.Token)
	err = a.redisCache.Get(cacheKey, &resetPassword)
	if err != nil {
		return nil, err
	}

	entUser, err := a.authRepository.FindByEmail(ctx, resetPassword.Email)
	if err != nil {
		return nil, err
	}

	// Update password
	entUser, err = a.authRepository.ChangePassword(ctx, entUser.ID.String(), req.NewPassword)
	if err != nil {
		return nil, err
	}

	// Delete token in redis
	err = a.redisCache.Delete(cacheKey)
	if err != nil {
		return nil, err
	}

	return &common.EmptyResp{
		Success: true,
	}, nil
}

// ChangePassword is a function that changes the password of a user.
func (a authService) ChangePassword(ctx context.Context, req *auth.ChangePasswordReq) (*common.EmptyResp, error) {
	principal := grpcutil.GetGrpcPrincipal(ctx)
	entUser, err := a.authRepository.FindById(ctx, principal.UserId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Validate request and check if old password is correct
	err = a.authValidator.ChangePassword(req, entUser.Password)
	if err != nil {
		return nil, err
	}

	// Save new password
	_, err = a.authRepository.ChangePassword(ctx, principal.UserId, req.NewPassword)
	if err != nil {
		return nil, err
	}

	return &common.EmptyResp{
		Success: true,
	}, nil
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
	kProducer := producer.NewKProducer(constants.TopicRegisterUser)
	go func() {
		registerUser := domain.RegisterUser{
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

	grpcUser := &grpcdomain.GrpcUser{
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
	redisCache *cache.RedisCache,
	authRepository authrepo.IAuthRepository,
	authValidator validator.IAuthValidator) IAuthService {
	return &authService{
		redisCache:     redisCache,
		authRepository: authRepository,
		authValidator:  authValidator,
	}
}
