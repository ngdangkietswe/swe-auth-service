package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/util"
)

// Action holds the schema definition for the Action entity.
type Action struct {
	ent.Schema
}

// Fields of the Action.
func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("description").Optional().Nillable(),
	}
}

// Edges of the Action.
func (Action) Edges() []ent.Edge {
	return []ent.Edge{
		util.One2Many("permissions", Permission.Type),
	}
}

// Annotations of the Action.
func (Action) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "action",
		},
	}
}
