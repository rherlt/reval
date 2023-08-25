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

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldName, v))
}

// ExternalId applies equality check predicate on the "externalId" field. It's identical to ExternalIdEQ.
func ExternalId(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldExternalId, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldDescription, v))
}

// Systemprompt applies equality check predicate on the "systemprompt" field. It's identical to SystempromptEQ.
func Systemprompt(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldSystemprompt, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldDate, v))
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

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContainsFold(FieldName, v))
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

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContainsFold(FieldDescription, v))
}

// SystempromptEQ applies the EQ predicate on the "systemprompt" field.
func SystempromptEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEQ(FieldSystemprompt, v))
}

// SystempromptNEQ applies the NEQ predicate on the "systemprompt" field.
func SystempromptNEQ(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNEQ(FieldSystemprompt, v))
}

// SystempromptIn applies the In predicate on the "systemprompt" field.
func SystempromptIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldIn(FieldSystemprompt, vs...))
}

// SystempromptNotIn applies the NotIn predicate on the "systemprompt" field.
func SystempromptNotIn(vs ...string) predicate.Scenario {
	return predicate.Scenario(sql.FieldNotIn(FieldSystemprompt, vs...))
}

// SystempromptGT applies the GT predicate on the "systemprompt" field.
func SystempromptGT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGT(FieldSystemprompt, v))
}

// SystempromptGTE applies the GTE predicate on the "systemprompt" field.
func SystempromptGTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldGTE(FieldSystemprompt, v))
}

// SystempromptLT applies the LT predicate on the "systemprompt" field.
func SystempromptLT(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLT(FieldSystemprompt, v))
}

// SystempromptLTE applies the LTE predicate on the "systemprompt" field.
func SystempromptLTE(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldLTE(FieldSystemprompt, v))
}

// SystempromptContains applies the Contains predicate on the "systemprompt" field.
func SystempromptContains(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContains(FieldSystemprompt, v))
}

// SystempromptHasPrefix applies the HasPrefix predicate on the "systemprompt" field.
func SystempromptHasPrefix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasPrefix(FieldSystemprompt, v))
}

// SystempromptHasSuffix applies the HasSuffix predicate on the "systemprompt" field.
func SystempromptHasSuffix(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldHasSuffix(FieldSystemprompt, v))
}

// SystempromptIsNil applies the IsNil predicate on the "systemprompt" field.
func SystempromptIsNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldIsNull(FieldSystemprompt))
}

// SystempromptNotNil applies the NotNil predicate on the "systemprompt" field.
func SystempromptNotNil() predicate.Scenario {
	return predicate.Scenario(sql.FieldNotNull(FieldSystemprompt))
}

// SystempromptEqualFold applies the EqualFold predicate on the "systemprompt" field.
func SystempromptEqualFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldEqualFold(FieldSystemprompt, v))
}

// SystempromptContainsFold applies the ContainsFold predicate on the "systemprompt" field.
func SystempromptContainsFold(v string) predicate.Scenario {
	return predicate.Scenario(sql.FieldContainsFold(FieldSystemprompt, v))
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
