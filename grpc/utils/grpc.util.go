package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ngdangkietswe/swe-auth-service/grpc/constant"
	"github.com/ngdangkietswe/swe-auth-service/grpc/domain"
)

// AsGrpcPrincipal is a function that converts a jwt.MapClaims to a SweGrpcPrincipal
func AsGrpcPrincipal(claims *jwt.MapClaims) (*domain.SweGrpcPrincipal, error) {
	claimsUser := (*claims)["user"].(map[string]interface{})
	bytes, err := json.Marshal(&claimsUser)

	if err != nil {
		return nil, fmt.Errorf("invalid principal in context")
	}

	principal := &domain.SweGrpcPrincipal{}
	err = json.Unmarshal(bytes, principal)
	return principal, err
}

// GetGrpcPrincipal is a function that gets the SweGrpcPrincipal from the context
func GetGrpcPrincipal(ctx context.Context) *domain.SweGrpcPrincipal {
	if principal, ok := ctx.Value(constant.CtxPrincipalKey).(*domain.SweGrpcPrincipal); ok {
		return principal
	}
	return nil
}
