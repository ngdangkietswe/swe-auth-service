package auth

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type IAuthValidator interface {
	RegisterUser(ctx context.Context, req *auth.User) error
}
