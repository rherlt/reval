package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type Response struct {
	ent.Schema
}

// Fields of the User.
func (Response) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("externalId").Optional(),
		field.UUID("requestId", uuid.UUID{}),
		field.String("from").Optional(),
		field.String("subject").Optional(),
		field.String("body"),
		field.Time("date").Optional(),
	}
}

// Edges of the User.
func (Response) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("request", Request.Type).
			Ref("responses").
			Unique().
			Required().
			Field("requestId"),
		//.Annotations(entsql.OnDelete(entsql.Restrict)),

		edge.To("evaluations", Evaluation.Type),
	}
}
