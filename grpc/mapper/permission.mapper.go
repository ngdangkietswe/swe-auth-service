package mapper

import (
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

// AsMonoPermission is a function that maps an ent.Permission to an auth.Permission.
func AsMonoPermission(entP *ent.Permission, entA *ent.Action, entR *ent.Resource) *auth.Permission {
	resp := &auth.Permission{
		Id:       entP.ID.String(),
		Action:   AsMonoAction(entA),
		Resource: AsMonoResource(entR),
	}

	if entP.Description != nil {
		resp.Description = *entP.Description
	}

	return resp
}

// AsListPermission is a function that maps a list of ent.Permission to a list of auth.Permission.
func AsListPermission(entPs []*ent.Permission, mapEntAction map[uuid.UUID]*ent.Action, mapEntResource map[uuid.UUID]*ent.Resource) []*auth.Permission {
	resp := make([]*auth.Permission, 0, len(entPs))
	for _, entP := range entPs {
		resp = append(resp, AsMonoPermission(entP, mapEntAction[entP.ActionID], mapEntResource[entP.ResourceID]))
	}
	return resp
}
