package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type Request struct {
	ent.Schema
}

// Fields of the User.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("externalId").Optional(),
		field.String("from").Optional(),
		field.String("subject").Optional(),
		field.String("body"),
		field.Time("date").Optional(),
	}
}

// Edges of the User.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("responses", Response.Type),
	}
}
