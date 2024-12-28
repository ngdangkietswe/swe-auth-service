package mapper

import (
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

// AsMonoAction is a function that maps an ent.Action to an auth.Action.
func AsMonoAction(entA *ent.Action) *auth.Action {
	resp := &auth.Action{
		Id:   entA.ID.String(),
		Name: entA.Name,
	}

	if entA.Description != nil {
		resp.Description = *entA.Description
	}

	return resp
}
