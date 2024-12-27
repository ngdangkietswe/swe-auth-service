package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UsersPermission holds the schema definition for the UsersPermission entity.
type UsersPermission struct {
	ent.Schema
}

// Fields of the UsersPermission.
func (UsersPermission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("permission_id", uuid.UUID{}),
	}
}

// Edges of the UsersPermission.
func (UsersPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("users_permissions"),
		edge.From("permission", Permission.Type).
			Ref("users_permissions"),
	}
}

// Annotations of the UsersPermission.
func (UsersPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "users_permission",
		},
	}
}
