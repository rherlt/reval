package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EvaluationPrompt holds the schema definition for the EvaluationPrompt entity.
type EvaluationPrompt struct {
	ent.Schema
}

// Fields of the EvaluationPrompt.
func (EvaluationPrompt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("prompt"),
	}
}

// Edges of the EvaluationPrompt.
func (EvaluationPrompt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("evaluations", Evaluation.Type),
	}
}
