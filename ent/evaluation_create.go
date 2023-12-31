// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rherlt/reval/ent/evaluation"
	"github.com/rherlt/reval/ent/response"
	"github.com/rherlt/reval/ent/user"
)

// EvaluationCreate is the builder for creating a Evaluation entity.
type EvaluationCreate struct {
	config
	mutation *EvaluationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserId sets the "userId" field.
func (ec *EvaluationCreate) SetUserId(u uuid.UUID) *EvaluationCreate {
	ec.mutation.SetUserId(u)
	return ec
}

// SetResponseId sets the "responseId" field.
func (ec *EvaluationCreate) SetResponseId(u uuid.UUID) *EvaluationCreate {
	ec.mutation.SetResponseId(u)
	return ec
}

// SetExternalId sets the "externalId" field.
func (ec *EvaluationCreate) SetExternalId(s string) *EvaluationCreate {
	ec.mutation.SetExternalId(s)
	return ec
}

// SetNillableExternalId sets the "externalId" field if the given value is not nil.
func (ec *EvaluationCreate) SetNillableExternalId(s *string) *EvaluationCreate {
	if s != nil {
		ec.SetExternalId(*s)
	}
	return ec
}

// SetDate sets the "date" field.
func (ec *EvaluationCreate) SetDate(t time.Time) *EvaluationCreate {
	ec.mutation.SetDate(t)
	return ec
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ec *EvaluationCreate) SetNillableDate(t *time.Time) *EvaluationCreate {
	if t != nil {
		ec.SetDate(*t)
	}
	return ec
}

// SetEvaluationResult sets the "evaluationResult" field.
func (ec *EvaluationCreate) SetEvaluationResult(s string) *EvaluationCreate {
	ec.mutation.SetEvaluationResult(s)
	return ec
}

// SetID sets the "id" field.
func (ec *EvaluationCreate) SetID(u uuid.UUID) *EvaluationCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EvaluationCreate) SetNillableID(u *uuid.UUID) *EvaluationCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ec *EvaluationCreate) SetUserID(id uuid.UUID) *EvaluationCreate {
	ec.mutation.SetUserID(id)
	return ec
}

// SetUser sets the "user" edge to the User entity.
func (ec *EvaluationCreate) SetUser(u *User) *EvaluationCreate {
	return ec.SetUserID(u.ID)
}

// SetResponseID sets the "response" edge to the Response entity by ID.
func (ec *EvaluationCreate) SetResponseID(id uuid.UUID) *EvaluationCreate {
	ec.mutation.SetResponseID(id)
	return ec
}

// SetResponse sets the "response" edge to the Response entity.
func (ec *EvaluationCreate) SetResponse(r *Response) *EvaluationCreate {
	return ec.SetResponseID(r.ID)
}

// Mutation returns the EvaluationMutation object of the builder.
func (ec *EvaluationCreate) Mutation() *EvaluationMutation {
	return ec.mutation
}

// Save creates the Evaluation in the database.
func (ec *EvaluationCreate) Save(ctx context.Context) (*Evaluation, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EvaluationCreate) SaveX(ctx context.Context) *Evaluation {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EvaluationCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EvaluationCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EvaluationCreate) defaults() {
	if _, ok := ec.mutation.ID(); !ok {
		v := evaluation.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EvaluationCreate) check() error {
	if _, ok := ec.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "Evaluation.userId"`)}
	}
	if _, ok := ec.mutation.ResponseId(); !ok {
		return &ValidationError{Name: "responseId", err: errors.New(`ent: missing required field "Evaluation.responseId"`)}
	}
	if _, ok := ec.mutation.EvaluationResult(); !ok {
		return &ValidationError{Name: "evaluationResult", err: errors.New(`ent: missing required field "Evaluation.evaluationResult"`)}
	}
	if _, ok := ec.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Evaluation.user"`)}
	}
	if _, ok := ec.mutation.ResponseID(); !ok {
		return &ValidationError{Name: "response", err: errors.New(`ent: missing required edge "Evaluation.response"`)}
	}
	return nil
}

func (ec *EvaluationCreate) sqlSave(ctx context.Context) (*Evaluation, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EvaluationCreate) createSpec() (*Evaluation, *sqlgraph.CreateSpec) {
	var (
		_node = &Evaluation{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(evaluation.Table, sqlgraph.NewFieldSpec(evaluation.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ec.conflict
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.ExternalId(); ok {
		_spec.SetField(evaluation.FieldExternalId, field.TypeString, value)
		_node.ExternalId = value
	}
	if value, ok := ec.mutation.Date(); ok {
		_spec.SetField(evaluation.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := ec.mutation.EvaluationResult(); ok {
		_spec.SetField(evaluation.FieldEvaluationResult, field.TypeString, value)
		_node.EvaluationResult = value
	}
	if nodes := ec.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   evaluation.UserTable,
			Columns: []string{evaluation.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   evaluation.ResponseTable,
			Columns: []string{evaluation.ResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(response.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ResponseId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Evaluation.Create().
//		SetUserId(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EvaluationUpsert) {
//			SetUserId(v+v).
//		}).
//		Exec(ctx)
func (ec *EvaluationCreate) OnConflict(opts ...sql.ConflictOption) *EvaluationUpsertOne {
	ec.conflict = opts
	return &EvaluationUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Evaluation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ec *EvaluationCreate) OnConflictColumns(columns ...string) *EvaluationUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EvaluationUpsertOne{
		create: ec,
	}
}

type (
	// EvaluationUpsertOne is the builder for "upsert"-ing
	//  one Evaluation node.
	EvaluationUpsertOne struct {
		create *EvaluationCreate
	}

	// EvaluationUpsert is the "OnConflict" setter.
	EvaluationUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserId sets the "userId" field.
func (u *EvaluationUpsert) SetUserId(v uuid.UUID) *EvaluationUpsert {
	u.Set(evaluation.FieldUserId, v)
	return u
}

// UpdateUserId sets the "userId" field to the value that was provided on create.
func (u *EvaluationUpsert) UpdateUserId() *EvaluationUpsert {
	u.SetExcluded(evaluation.FieldUserId)
	return u
}

// SetResponseId sets the "responseId" field.
func (u *EvaluationUpsert) SetResponseId(v uuid.UUID) *EvaluationUpsert {
	u.Set(evaluation.FieldResponseId, v)
	return u
}

// UpdateResponseId sets the "responseId" field to the value that was provided on create.
func (u *EvaluationUpsert) UpdateResponseId() *EvaluationUpsert {
	u.SetExcluded(evaluation.FieldResponseId)
	return u
}

// SetExternalId sets the "externalId" field.
func (u *EvaluationUpsert) SetExternalId(v string) *EvaluationUpsert {
	u.Set(evaluation.FieldExternalId, v)
	return u
}

// UpdateExternalId sets the "externalId" field to the value that was provided on create.
func (u *EvaluationUpsert) UpdateExternalId() *EvaluationUpsert {
	u.SetExcluded(evaluation.FieldExternalId)
	return u
}

// ClearExternalId clears the value of the "externalId" field.
func (u *EvaluationUpsert) ClearExternalId() *EvaluationUpsert {
	u.SetNull(evaluation.FieldExternalId)
	return u
}

// SetDate sets the "date" field.
func (u *EvaluationUpsert) SetDate(v time.Time) *EvaluationUpsert {
	u.Set(evaluation.FieldDate, v)
	return u
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *EvaluationUpsert) UpdateDate() *EvaluationUpsert {
	u.SetExcluded(evaluation.FieldDate)
	return u
}

// ClearDate clears the value of the "date" field.
func (u *EvaluationUpsert) ClearDate() *EvaluationUpsert {
	u.SetNull(evaluation.FieldDate)
	return u
}

// SetEvaluationResult sets the "evaluationResult" field.
func (u *EvaluationUpsert) SetEvaluationResult(v string) *EvaluationUpsert {
	u.Set(evaluation.FieldEvaluationResult, v)
	return u
}

// UpdateEvaluationResult sets the "evaluationResult" field to the value that was provided on create.
func (u *EvaluationUpsert) UpdateEvaluationResult() *EvaluationUpsert {
	u.SetExcluded(evaluation.FieldEvaluationResult)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Evaluation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(evaluation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EvaluationUpsertOne) UpdateNewValues() *EvaluationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(evaluation.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Evaluation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EvaluationUpsertOne) Ignore() *EvaluationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EvaluationUpsertOne) DoNothing() *EvaluationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EvaluationCreate.OnConflict
// documentation for more info.
func (u *EvaluationUpsertOne) Update(set func(*EvaluationUpsert)) *EvaluationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EvaluationUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserId sets the "userId" field.
func (u *EvaluationUpsertOne) SetUserId(v uuid.UUID) *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetUserId(v)
	})
}

// UpdateUserId sets the "userId" field to the value that was provided on create.
func (u *EvaluationUpsertOne) UpdateUserId() *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateUserId()
	})
}

// SetResponseId sets the "responseId" field.
func (u *EvaluationUpsertOne) SetResponseId(v uuid.UUID) *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetResponseId(v)
	})
}

// UpdateResponseId sets the "responseId" field to the value that was provided on create.
func (u *EvaluationUpsertOne) UpdateResponseId() *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateResponseId()
	})
}

// SetExternalId sets the "externalId" field.
func (u *EvaluationUpsertOne) SetExternalId(v string) *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetExternalId(v)
	})
}

// UpdateExternalId sets the "externalId" field to the value that was provided on create.
func (u *EvaluationUpsertOne) UpdateExternalId() *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateExternalId()
	})
}

// ClearExternalId clears the value of the "externalId" field.
func (u *EvaluationUpsertOne) ClearExternalId() *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.ClearExternalId()
	})
}

// SetDate sets the "date" field.
func (u *EvaluationUpsertOne) SetDate(v time.Time) *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetDate(v)
	})
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *EvaluationUpsertOne) UpdateDate() *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateDate()
	})
}

// ClearDate clears the value of the "date" field.
func (u *EvaluationUpsertOne) ClearDate() *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.ClearDate()
	})
}

// SetEvaluationResult sets the "evaluationResult" field.
func (u *EvaluationUpsertOne) SetEvaluationResult(v string) *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetEvaluationResult(v)
	})
}

// UpdateEvaluationResult sets the "evaluationResult" field to the value that was provided on create.
func (u *EvaluationUpsertOne) UpdateEvaluationResult() *EvaluationUpsertOne {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateEvaluationResult()
	})
}

// Exec executes the query.
func (u *EvaluationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EvaluationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EvaluationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EvaluationUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: EvaluationUpsertOne.ID is not supported by MySQL driver. Use EvaluationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EvaluationUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EvaluationCreateBulk is the builder for creating many Evaluation entities in bulk.
type EvaluationCreateBulk struct {
	config
	builders []*EvaluationCreate
	conflict []sql.ConflictOption
}

// Save creates the Evaluation entities in the database.
func (ecb *EvaluationCreateBulk) Save(ctx context.Context) ([]*Evaluation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Evaluation, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EvaluationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EvaluationCreateBulk) SaveX(ctx context.Context) []*Evaluation {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EvaluationCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EvaluationCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Evaluation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EvaluationUpsert) {
//			SetUserId(v+v).
//		}).
//		Exec(ctx)
func (ecb *EvaluationCreateBulk) OnConflict(opts ...sql.ConflictOption) *EvaluationUpsertBulk {
	ecb.conflict = opts
	return &EvaluationUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Evaluation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecb *EvaluationCreateBulk) OnConflictColumns(columns ...string) *EvaluationUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EvaluationUpsertBulk{
		create: ecb,
	}
}

// EvaluationUpsertBulk is the builder for "upsert"-ing
// a bulk of Evaluation nodes.
type EvaluationUpsertBulk struct {
	create *EvaluationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Evaluation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(evaluation.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EvaluationUpsertBulk) UpdateNewValues() *EvaluationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(evaluation.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Evaluation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EvaluationUpsertBulk) Ignore() *EvaluationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EvaluationUpsertBulk) DoNothing() *EvaluationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EvaluationCreateBulk.OnConflict
// documentation for more info.
func (u *EvaluationUpsertBulk) Update(set func(*EvaluationUpsert)) *EvaluationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EvaluationUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserId sets the "userId" field.
func (u *EvaluationUpsertBulk) SetUserId(v uuid.UUID) *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetUserId(v)
	})
}

// UpdateUserId sets the "userId" field to the value that was provided on create.
func (u *EvaluationUpsertBulk) UpdateUserId() *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateUserId()
	})
}

// SetResponseId sets the "responseId" field.
func (u *EvaluationUpsertBulk) SetResponseId(v uuid.UUID) *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetResponseId(v)
	})
}

// UpdateResponseId sets the "responseId" field to the value that was provided on create.
func (u *EvaluationUpsertBulk) UpdateResponseId() *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateResponseId()
	})
}

// SetExternalId sets the "externalId" field.
func (u *EvaluationUpsertBulk) SetExternalId(v string) *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetExternalId(v)
	})
}

// UpdateExternalId sets the "externalId" field to the value that was provided on create.
func (u *EvaluationUpsertBulk) UpdateExternalId() *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateExternalId()
	})
}

// ClearExternalId clears the value of the "externalId" field.
func (u *EvaluationUpsertBulk) ClearExternalId() *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.ClearExternalId()
	})
}

// SetDate sets the "date" field.
func (u *EvaluationUpsertBulk) SetDate(v time.Time) *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetDate(v)
	})
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *EvaluationUpsertBulk) UpdateDate() *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateDate()
	})
}

// ClearDate clears the value of the "date" field.
func (u *EvaluationUpsertBulk) ClearDate() *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.ClearDate()
	})
}

// SetEvaluationResult sets the "evaluationResult" field.
func (u *EvaluationUpsertBulk) SetEvaluationResult(v string) *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.SetEvaluationResult(v)
	})
}

// UpdateEvaluationResult sets the "evaluationResult" field to the value that was provided on create.
func (u *EvaluationUpsertBulk) UpdateEvaluationResult() *EvaluationUpsertBulk {
	return u.Update(func(s *EvaluationUpsert) {
		s.UpdateEvaluationResult()
	})
}

// Exec executes the query.
func (u *EvaluationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EvaluationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EvaluationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EvaluationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
