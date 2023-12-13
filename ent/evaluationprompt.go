// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rherlt/reval/ent/evaluationprompt"
)

// EvaluationPrompt is the model entity for the EvaluationPrompt schema.
type EvaluationPrompt struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Prompt holds the value of the "prompt" field.
	Prompt string `json:"prompt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EvaluationPromptQuery when eager-loading is set.
	Edges        EvaluationPromptEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EvaluationPromptEdges holds the relations/edges for other nodes in the graph.
type EvaluationPromptEdges struct {
	// Evaluations holds the value of the evaluations edge.
	Evaluations []*Evaluation `json:"evaluations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EvaluationsOrErr returns the Evaluations value or an error if the edge
// was not loaded in eager-loading.
func (e EvaluationPromptEdges) EvaluationsOrErr() ([]*Evaluation, error) {
	if e.loadedTypes[0] {
		return e.Evaluations, nil
	}
	return nil, &NotLoadedError{edge: "evaluations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EvaluationPrompt) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case evaluationprompt.FieldPrompt:
			values[i] = new(sql.NullString)
		case evaluationprompt.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EvaluationPrompt fields.
func (ep *EvaluationPrompt) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case evaluationprompt.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ep.ID = *value
			}
		case evaluationprompt.FieldPrompt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prompt", values[i])
			} else if value.Valid {
				ep.Prompt = value.String
			}
		default:
			ep.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EvaluationPrompt.
// This includes values selected through modifiers, order, etc.
func (ep *EvaluationPrompt) Value(name string) (ent.Value, error) {
	return ep.selectValues.Get(name)
}

// QueryEvaluations queries the "evaluations" edge of the EvaluationPrompt entity.
func (ep *EvaluationPrompt) QueryEvaluations() *EvaluationQuery {
	return NewEvaluationPromptClient(ep.config).QueryEvaluations(ep)
}

// Update returns a builder for updating this EvaluationPrompt.
// Note that you need to call EvaluationPrompt.Unwrap() before calling this method if this EvaluationPrompt
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *EvaluationPrompt) Update() *EvaluationPromptUpdateOne {
	return NewEvaluationPromptClient(ep.config).UpdateOne(ep)
}

// Unwrap unwraps the EvaluationPrompt entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ep *EvaluationPrompt) Unwrap() *EvaluationPrompt {
	_tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("ent: EvaluationPrompt is not a transactional entity")
	}
	ep.config.driver = _tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *EvaluationPrompt) String() string {
	var builder strings.Builder
	builder.WriteString("EvaluationPrompt(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ep.ID))
	builder.WriteString("prompt=")
	builder.WriteString(ep.Prompt)
	builder.WriteByte(')')
	return builder.String()
}

// EvaluationPrompts is a parsable slice of EvaluationPrompt.
type EvaluationPrompts []*EvaluationPrompt