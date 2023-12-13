// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rherlt/reval/ent/evaluationprompt"
	"github.com/rherlt/reval/ent/predicate"
)

// EvaluationPromptDelete is the builder for deleting a EvaluationPrompt entity.
type EvaluationPromptDelete struct {
	config
	hooks    []Hook
	mutation *EvaluationPromptMutation
}

// Where appends a list predicates to the EvaluationPromptDelete builder.
func (epd *EvaluationPromptDelete) Where(ps ...predicate.EvaluationPrompt) *EvaluationPromptDelete {
	epd.mutation.Where(ps...)
	return epd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (epd *EvaluationPromptDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, epd.sqlExec, epd.mutation, epd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (epd *EvaluationPromptDelete) ExecX(ctx context.Context) int {
	n, err := epd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (epd *EvaluationPromptDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(evaluationprompt.Table, sqlgraph.NewFieldSpec(evaluationprompt.FieldID, field.TypeUUID))
	if ps := epd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, epd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	epd.mutation.done = true
	return affected, err
}

// EvaluationPromptDeleteOne is the builder for deleting a single EvaluationPrompt entity.
type EvaluationPromptDeleteOne struct {
	epd *EvaluationPromptDelete
}

// Where appends a list predicates to the EvaluationPromptDelete builder.
func (epdo *EvaluationPromptDeleteOne) Where(ps ...predicate.EvaluationPrompt) *EvaluationPromptDeleteOne {
	epdo.epd.mutation.Where(ps...)
	return epdo
}

// Exec executes the deletion query.
func (epdo *EvaluationPromptDeleteOne) Exec(ctx context.Context) error {
	n, err := epdo.epd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{evaluationprompt.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (epdo *EvaluationPromptDeleteOne) ExecX(ctx context.Context) {
	if err := epdo.Exec(ctx); err != nil {
		panic(err)
	}
}