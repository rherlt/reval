package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type Evaluation struct {
	ent.Schema
}

// Fields of the User.
func (Evaluation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("userId", uuid.UUID{}),
		field.UUID("responseId", uuid.UUID{}),
		field.String("externalId").Optional(),
		field.Time("date"),
		field.String("evaluationResult"),
	}
}

// Edges of the User.
func (Evaluation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("evaluations").
			Unique().
			Required().
			Field("userId"),
		//Annotations(entsql.OnDelete(entsql.Restrict)),

		edge.From("response", Response.Type).
			Ref("evaluations").
			Unique().
			Required().
			Field("responseId"),
		//Annotations(entsql.OnDelete(entsql.Restrict)),
	}
}
