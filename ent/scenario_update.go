// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rherlt/reval/ent/predicate"
	"github.com/rherlt/reval/ent/response"
	"github.com/rherlt/reval/ent/scenario"
)

// ScenarioUpdate is the builder for updating Scenario entities.
type ScenarioUpdate struct {
	config
	hooks    []Hook
	mutation *ScenarioMutation
}

// Where appends a list predicates to the ScenarioUpdate builder.
func (su *ScenarioUpdate) Where(ps ...predicate.Scenario) *ScenarioUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *ScenarioUpdate) SetName(s string) *ScenarioUpdate {
	su.mutation.SetName(s)
	return su
}

// SetExternalId sets the "externalId" field.
func (su *ScenarioUpdate) SetExternalId(s string) *ScenarioUpdate {
	su.mutation.SetExternalId(s)
	return su
}

// SetNillableExternalId sets the "externalId" field if the given value is not nil.
func (su *ScenarioUpdate) SetNillableExternalId(s *string) *ScenarioUpdate {
	if s != nil {
		su.SetExternalId(*s)
	}
	return su
}

// ClearExternalId clears the value of the "externalId" field.
func (su *ScenarioUpdate) ClearExternalId() *ScenarioUpdate {
	su.mutation.ClearExternalId()
	return su
}

// SetDesctiption sets the "desctiption" field.
func (su *ScenarioUpdate) SetDesctiption(s string) *ScenarioUpdate {
	su.mutation.SetDesctiption(s)
	return su
}

// SetNillableDesctiption sets the "desctiption" field if the given value is not nil.
func (su *ScenarioUpdate) SetNillableDesctiption(s *string) *ScenarioUpdate {
	if s != nil {
		su.SetDesctiption(*s)
	}
	return su
}

// ClearDesctiption clears the value of the "desctiption" field.
func (su *ScenarioUpdate) ClearDesctiption() *ScenarioUpdate {
	su.mutation.ClearDesctiption()
	return su
}

// SetDate sets the "date" field.
func (su *ScenarioUpdate) SetDate(t time.Time) *ScenarioUpdate {
	su.mutation.SetDate(t)
	return su
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (su *ScenarioUpdate) SetNillableDate(t *time.Time) *ScenarioUpdate {
	if t != nil {
		su.SetDate(*t)
	}
	return su
}

// ClearDate clears the value of the "date" field.
func (su *ScenarioUpdate) ClearDate() *ScenarioUpdate {
	su.mutation.ClearDate()
	return su
}

// AddResponseIDs adds the "responses" edge to the Response entity by IDs.
func (su *ScenarioUpdate) AddResponseIDs(ids ...uuid.UUID) *ScenarioUpdate {
	su.mutation.AddResponseIDs(ids...)
	return su
}

// AddResponses adds the "responses" edges to the Response entity.
func (su *ScenarioUpdate) AddResponses(r ...*Response) *ScenarioUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return su.AddResponseIDs(ids...)
}

// Mutation returns the ScenarioMutation object of the builder.
func (su *ScenarioUpdate) Mutation() *ScenarioMutation {
	return su.mutation
}

// ClearResponses clears all "responses" edges to the Response entity.
func (su *ScenarioUpdate) ClearResponses() *ScenarioUpdate {
	su.mutation.ClearResponses()
	return su
}

// RemoveResponseIDs removes the "responses" edge to Response entities by IDs.
func (su *ScenarioUpdate) RemoveResponseIDs(ids ...uuid.UUID) *ScenarioUpdate {
	su.mutation.RemoveResponseIDs(ids...)
	return su
}

// RemoveResponses removes "responses" edges to Response entities.
func (su *ScenarioUpdate) RemoveResponses(r ...*Response) *ScenarioUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return su.RemoveResponseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ScenarioUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ScenarioUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ScenarioUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ScenarioUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ScenarioUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(scenario.Table, scenario.Columns, sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(scenario.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.ExternalId(); ok {
		_spec.SetField(scenario.FieldExternalId, field.TypeString, value)
	}
	if su.mutation.ExternalIdCleared() {
		_spec.ClearField(scenario.FieldExternalId, field.TypeString)
	}
	if value, ok := su.mutation.Desctiption(); ok {
		_spec.SetField(scenario.FieldDesctiption, field.TypeString, value)
	}
	if su.mutation.DesctiptionCleared() {
		_spec.ClearField(scenario.FieldDesctiption, field.TypeString)
	}
	if value, ok := su.mutation.Date(); ok {
		_spec.SetField(scenario.FieldDate, field.TypeTime, value)
	}
	if su.mutation.DateCleared() {
		_spec.ClearField(scenario.FieldDate, field.TypeTime)
	}
	if su.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenario.ResponsesTable,
			Columns: []string{scenario.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedResponsesIDs(); len(nodes) > 0 && !su.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenario.ResponsesTable,
			Columns: []string{scenario.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenario.ResponsesTable,
			Columns: []string{scenario.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scenario.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ScenarioUpdateOne is the builder for updating a single Scenario entity.
type ScenarioUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScenarioMutation
}

// SetName sets the "name" field.
func (suo *ScenarioUpdateOne) SetName(s string) *ScenarioUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetExternalId sets the "externalId" field.
func (suo *ScenarioUpdateOne) SetExternalId(s string) *ScenarioUpdateOne {
	suo.mutation.SetExternalId(s)
	return suo
}

// SetNillableExternalId sets the "externalId" field if the given value is not nil.
func (suo *ScenarioUpdateOne) SetNillableExternalId(s *string) *ScenarioUpdateOne {
	if s != nil {
		suo.SetExternalId(*s)
	}
	return suo
}

// ClearExternalId clears the value of the "externalId" field.
func (suo *ScenarioUpdateOne) ClearExternalId() *ScenarioUpdateOne {
	suo.mutation.ClearExternalId()
	return suo
}

// SetDesctiption sets the "desctiption" field.
func (suo *ScenarioUpdateOne) SetDesctiption(s string) *ScenarioUpdateOne {
	suo.mutation.SetDesctiption(s)
	return suo
}

// SetNillableDesctiption sets the "desctiption" field if the given value is not nil.
func (suo *ScenarioUpdateOne) SetNillableDesctiption(s *string) *ScenarioUpdateOne {
	if s != nil {
		suo.SetDesctiption(*s)
	}
	return suo
}

// ClearDesctiption clears the value of the "desctiption" field.
func (suo *ScenarioUpdateOne) ClearDesctiption() *ScenarioUpdateOne {
	suo.mutation.ClearDesctiption()
	return suo
}

// SetDate sets the "date" field.
func (suo *ScenarioUpdateOne) SetDate(t time.Time) *ScenarioUpdateOne {
	suo.mutation.SetDate(t)
	return suo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (suo *ScenarioUpdateOne) SetNillableDate(t *time.Time) *ScenarioUpdateOne {
	if t != nil {
		suo.SetDate(*t)
	}
	return suo
}

// ClearDate clears the value of the "date" field.
func (suo *ScenarioUpdateOne) ClearDate() *ScenarioUpdateOne {
	suo.mutation.ClearDate()
	return suo
}

// AddResponseIDs adds the "responses" edge to the Response entity by IDs.
func (suo *ScenarioUpdateOne) AddResponseIDs(ids ...uuid.UUID) *ScenarioUpdateOne {
	suo.mutation.AddResponseIDs(ids...)
	return suo
}

// AddResponses adds the "responses" edges to the Response entity.
func (suo *ScenarioUpdateOne) AddResponses(r ...*Response) *ScenarioUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return suo.AddResponseIDs(ids...)
}

// Mutation returns the ScenarioMutation object of the builder.
func (suo *ScenarioUpdateOne) Mutation() *ScenarioMutation {
	return suo.mutation
}

// ClearResponses clears all "responses" edges to the Response entity.
func (suo *ScenarioUpdateOne) ClearResponses() *ScenarioUpdateOne {
	suo.mutation.ClearResponses()
	return suo
}

// RemoveResponseIDs removes the "responses" edge to Response entities by IDs.
func (suo *ScenarioUpdateOne) RemoveResponseIDs(ids ...uuid.UUID) *ScenarioUpdateOne {
	suo.mutation.RemoveResponseIDs(ids...)
	return suo
}

// RemoveResponses removes "responses" edges to Response entities.
func (suo *ScenarioUpdateOne) RemoveResponses(r ...*Response) *ScenarioUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return suo.RemoveResponseIDs(ids...)
}

// Where appends a list predicates to the ScenarioUpdate builder.
func (suo *ScenarioUpdateOne) Where(ps ...predicate.Scenario) *ScenarioUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ScenarioUpdateOne) Select(field string, fields ...string) *ScenarioUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Scenario entity.
func (suo *ScenarioUpdateOne) Save(ctx context.Context) (*Scenario, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ScenarioUpdateOne) SaveX(ctx context.Context) *Scenario {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ScenarioUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ScenarioUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ScenarioUpdateOne) sqlSave(ctx context.Context) (_node *Scenario, err error) {
	_spec := sqlgraph.NewUpdateSpec(scenario.Table, scenario.Columns, sqlgraph.NewFieldSpec(scenario.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Scenario.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scenario.FieldID)
		for _, f := range fields {
			if !scenario.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != scenario.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(scenario.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.ExternalId(); ok {
		_spec.SetField(scenario.FieldExternalId, field.TypeString, value)
	}
	if suo.mutation.ExternalIdCleared() {
		_spec.ClearField(scenario.FieldExternalId, field.TypeString)
	}
	if value, ok := suo.mutation.Desctiption(); ok {
		_spec.SetField(scenario.FieldDesctiption, field.TypeString, value)
	}
	if suo.mutation.DesctiptionCleared() {
		_spec.ClearField(scenario.FieldDesctiption, field.TypeString)
	}
	if value, ok := suo.mutation.Date(); ok {
		_spec.SetField(scenario.FieldDate, field.TypeTime, value)
	}
	if suo.mutation.DateCleared() {
		_spec.ClearField(scenario.FieldDate, field.TypeTime)
	}
	if suo.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenario.ResponsesTable,
			Columns: []string{scenario.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedResponsesIDs(); len(nodes) > 0 && !suo.mutation.ResponsesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenario.ResponsesTable,
			Columns: []string{scenario.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ResponsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scenario.ResponsesTable,
			Columns: []string{scenario.ResponsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Scenario{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scenario.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}