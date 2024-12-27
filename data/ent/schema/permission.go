package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("action_id", uuid.UUID{}),
		field.UUID("resource_id", uuid.UUID{}),
		field.String("description").Optional().Nillable(),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("action", Action.Type).
			Ref("permissions"),
		edge.From("resource", Resource.Type).
			Ref("permissions"),
		edge.To("users_permissions", UsersPermission.Type),
	}
}

// Annotations of the Permission.
func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "permission",
		},
	}
}
