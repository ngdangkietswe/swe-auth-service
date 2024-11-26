package mapper

import (
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

func AsFailed(message string) (*auth.LoginResp, error) {
	return &auth.LoginResp{
		Success: false,
		Resp: &auth.LoginResp_Error{
			Error: &common.Error{
				Code:    401,
				Message: message,
			},
		},
	}, nil
}
