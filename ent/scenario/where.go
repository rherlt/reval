// Code generated by ent, DO NOT EDIT.

package scenario

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/rherlt/reval/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Scenario {
	return predicate.Scenario(sql.FieldLTE(FieldID, id))
}

// ExternalId applies equality check predicate on the "externalId" field. It's identical to ExternalIdEQ.
func ExternalId(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldExternalId, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldName, v))
}

// Desctiption applies equality check predicate on the "desctiption" field. It's identical to DesctiptionEQ.
func Desctiption(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldDesctiption, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldDate, v))
}

// ExternalIdEQ applies the EQ predicate on the "externalId" field.
func ExternalIdEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldExternalId, v))
}

// ExternalIdNEQ applies the NEQ predicate on the "externalId" field.
func ExternalIdNEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNEQ(FieldExternalId, v))
}

// ExternalIdIn applies the In predicate on the "externalId" field.
func ExternalIdIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldIn(FieldExternalId, vs...))
}

// ExternalIdNotIn applies the NotIn predicate on the "externalId" field.
func ExternalIdNotIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNotIn(FieldExternalId, vs...))
}

// ExternalIdGT applies the GT predicate on the "externalId" field.
func ExternalIdGT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGT(FieldExternalId, v))
}

// ExternalIdGTE applies the GTE predicate on the "externalId" field.
func ExternalIdGTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGTE(FieldExternalId, v))
}

// ExternalIdLT applies the LT predicate on the "externalId" field.
func ExternalIdLT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLT(FieldExternalId, v))
}

// ExternalIdLTE applies the LTE predicate on the "externalId" field.
func ExternalIdLTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLTE(FieldExternalId, v))
}

// ExternalIdContains applies the Contains predicate on the "externalId" field.
func ExternalIdContains(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContains(FieldExternalId, v))
}

// ExternalIdHasPrefix applies the HasPrefix predicate on the "externalId" field.
func ExternalIdHasPrefix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasPrefix(FieldExternalId, v))
}

// ExternalIdHasSuffix applies the HasSuffix predicate on the "externalId" field.
func ExternalIdHasSuffix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasSuffix(FieldExternalId, v))
}

// ExternalIdIsNil applies the IsNil predicate on the "externalId" field.
func ExternalIdIsNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldIsNull(FieldExternalId))
}

// ExternalIdNotNil applies the NotNil predicate on the "externalId" field.
func ExternalIdNotNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldNotNull(FieldExternalId))
}

// ExternalIdEqualFold applies the EqualFold predicate on the "externalId" field.
func ExternalIdEqualFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEqualFold(FieldExternalId, v))
}

// ExternalIdContainsFold applies the ContainsFold predicate on the "externalId" field.
func ExternalIdContainsFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContainsFold(FieldExternalId, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContainsFold(FieldName, v))
}

// DesctiptionEQ applies the EQ predicate on the "desctiption" field.
func DesctiptionEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldDesctiption, v))
}

// DesctiptionNEQ applies the NEQ predicate on the "desctiption" field.
func DesctiptionNEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNEQ(FieldDesctiption, v))
}

// DesctiptionIn applies the In predicate on the "desctiption" field.
func DesctiptionIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldIn(FieldDesctiption, vs...))
}

// DesctiptionNotIn applies the NotIn predicate on the "desctiption" field.
func DesctiptionNotIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNotIn(FieldDesctiption, vs...))
}

// DesctiptionGT applies the GT predicate on the "desctiption" field.
func DesctiptionGT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGT(FieldDesctiption, v))
}

// DesctiptionGTE applies the GTE predicate on the "desctiption" field.
func DesctiptionGTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGTE(FieldDesctiption, v))
}

// DesctiptionLT applies the LT predicate on the "desctiption" field.
func DesctiptionLT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLT(FieldDesctiption, v))
}

// DesctiptionLTE applies the LTE predicate on the "desctiption" field.
func DesctiptionLTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLTE(FieldDesctiption, v))
}

// DesctiptionContains applies the Contains predicate on the "desctiption" field.
func DesctiptionContains(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContains(FieldDesctiption, v))
}

// DesctiptionHasPrefix applies the HasPrefix predicate on the "desctiption" field.
func DesctiptionHasPrefix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasPrefix(FieldDesctiption, v))
}

// DesctiptionHasSuffix applies the HasSuffix predicate on the "desctiption" field.
func DesctiptionHasSuffix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasSuffix(FieldDesctiption, v))
}

// DesctiptionIsNil applies the IsNil predicate on the "desctiption" field.
func DesctiptionIsNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldIsNull(FieldDesctiption))
}

// DesctiptionNotNil applies the NotNil predicate on the "desctiption" field.
func DesctiptionNotNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldNotNull(FieldDesctiption))
}

// DesctiptionEqualFold applies the EqualFold predicate on the "desctiption" field.
func DesctiptionEqualFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEqualFold(FieldDesctiption, v))
}

// DesctiptionContainsFold applies the ContainsFold predicate on the "desctiption" field.
func DesctiptionContainsFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContainsFold(FieldDesctiption, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldLTE(FieldDate, v))
}

// DateIsNil applies the IsNil predicate on the "date" field.
func DateIsNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldIsNull(FieldDate))
}

// DateNotNil applies the NotNil predicate on the "date" field.
func DateNotNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldNotNull(FieldDate))
}

// HasResponses applies the HasEdge predicate on the "responses" edge.
func HasResponses() predicate.Scenario {
	return predicate.Scenario(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResponsesTable, ResponsesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResponsesWith applies the HasEdge predicate on the "responses" edge with a given conditions (other predicates).
func HasResponsesWith(preds ...predicate.Response) predicate.Scenario {
	return predicate.Scenario(func(s *sql.Selector) {
		step := newResponsesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Scenario) predicate.Scenario {
	return predicate.Scenario(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Scenario) predicate.Scenario {
	return predicate.Scenario(func(s *sql.Selector) {
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
func Not(p predicate.Scenario) predicate.Scenario {
	return predicate.Scenario(func(s *sql.Selector) {
		p(s.Not())
	})
}
