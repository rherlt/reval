// Code generated by ent, DO NOT EDIT.

package evaluationprompt

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/rherlt/reval/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldLTE(FieldID, id))
}

// Prompt applies equality check predicate on the "prompt" field. It's identical to PromptEQ.
func Prompt(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldEQ(FieldPrompt, v))
}

// PromptEQ applies the EQ predicate on the "prompt" field.
func PromptEQ(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldEQ(FieldPrompt, v))
}

// PromptNEQ applies the NEQ predicate on the "prompt" field.
func PromptNEQ(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldNEQ(FieldPrompt, v))
}

// PromptIn applies the In predicate on the "prompt" field.
func PromptIn(vs ...string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldIn(FieldPrompt, vs...))
}

// PromptNotIn applies the NotIn predicate on the "prompt" field.
func PromptNotIn(vs ...string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldNotIn(FieldPrompt, vs...))
}

// PromptGT applies the GT predicate on the "prompt" field.
func PromptGT(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldGT(FieldPrompt, v))
}

// PromptGTE applies the GTE predicate on the "prompt" field.
func PromptGTE(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldGTE(FieldPrompt, v))
}

// PromptLT applies the LT predicate on the "prompt" field.
func PromptLT(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldLT(FieldPrompt, v))
}

// PromptLTE applies the LTE predicate on the "prompt" field.
func PromptLTE(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldLTE(FieldPrompt, v))
}

// PromptContains applies the Contains predicate on the "prompt" field.
func PromptContains(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldContains(FieldPrompt, v))
}

// PromptHasPrefix applies the HasPrefix predicate on the "prompt" field.
func PromptHasPrefix(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldHasPrefix(FieldPrompt, v))
}

// PromptHasSuffix applies the HasSuffix predicate on the "prompt" field.
func PromptHasSuffix(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldHasSuffix(FieldPrompt, v))
}

// PromptEqualFold applies the EqualFold predicate on the "prompt" field.
func PromptEqualFold(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldEqualFold(FieldPrompt, v))
}

// PromptContainsFold applies the ContainsFold predicate on the "prompt" field.
func PromptContainsFold(v string) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(sql.FieldContainsFold(FieldPrompt, v))
}

// HasEvaluations applies the HasEdge predicate on the "evaluations" edge.
func HasEvaluations() predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EvaluationsTable, EvaluationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEvaluationsWith applies the HasEdge predicate on the "evaluations" edge with a given conditions (other predicates).
func HasEvaluationsWith(preds ...predicate.Evaluation) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(func(s *sql.Selector) {
		step := newEvaluationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EvaluationPrompt) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EvaluationPrompt) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EvaluationPrompt) predicate.EvaluationPrompt {
	return predicate.EvaluationPrompt(func(s *sql.Selector) {
		p(s.Not())
	})
}