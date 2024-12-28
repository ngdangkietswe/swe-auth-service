package mapper

import (
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

// AsMonoResource is a function that maps an ent.Resource to an auth.Resource.
func AsMonoResource(entR *ent.Resource) *auth.Resource {
	resp := &auth.Resource{
		Id:   entR.ID.String(),
		Name: entR.Name,
	}

	if entR.Description != nil {
		resp.Description = *entR.Description
	}

	return resp
}
