// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rherlt/reval/ent/evaluation"
	"github.com/rherlt/reval/ent/predicate"
	"github.com/rherlt/reval/ent/response"
	"github.com/rherlt/reval/ent/user"
)

// EvaluationQuery is the builder for querying Evaluation entities.
type EvaluationQuery struct {
	config
	ctx          *QueryContext
	order        []evaluation.OrderOption
	inters       []Interceptor
	predicates   []predicate.Evaluation
	withUser     *UserQuery
	withResponse *ResponseQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EvaluationQuery builder.
func (eq *EvaluationQuery) Where(ps ...predicate.Evaluation) *EvaluationQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EvaluationQuery) Limit(limit int) *EvaluationQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EvaluationQuery) Offset(offset int) *EvaluationQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EvaluationQuery) Unique(unique bool) *EvaluationQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EvaluationQuery) Order(o ...evaluation.OrderOption) *EvaluationQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryUser chains the current query on the "user" edge.
func (eq *EvaluationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(evaluation.Table, evaluation.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, evaluation.UserTable, evaluation.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResponse chains the current query on the "response" edge.
func (eq *EvaluationQuery) QueryResponse() *ResponseQuery {
	query := (&ResponseClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(evaluation.Table, evaluation.FieldID, selector),
			sqlgraph.To(response.Table, response.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, evaluation.ResponseTable, evaluation.ResponseColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Evaluation entity from the query.
// Returns a *NotFoundError when no Evaluation was found.
func (eq *EvaluationQuery) First(ctx context.Context) (*Evaluation, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{evaluation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EvaluationQuery) FirstX(ctx context.Context) *Evaluation {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Evaluation ID from the query.
// Returns a *NotFoundError when no Evaluation ID was found.
func (eq *EvaluationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{evaluation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EvaluationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Evaluation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Evaluation entity is found.
// Returns a *NotFoundError when no Evaluation entities are found.
func (eq *EvaluationQuery) Only(ctx context.Context) (*Evaluation, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{evaluation.Label}
	default:
		return nil, &NotSingularError{evaluation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EvaluationQuery) OnlyX(ctx context.Context) *Evaluation {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Evaluation ID in the query.
// Returns a *NotSingularError when more than one Evaluation ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EvaluationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{evaluation.Label}
	default:
		err = &NotSingularError{evaluation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EvaluationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Evaluations.
func (eq *EvaluationQuery) All(ctx context.Context) ([]*Evaluation, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Evaluation, *EvaluationQuery]()
	return withInterceptors[[]*Evaluation](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EvaluationQuery) AllX(ctx context.Context) []*Evaluation {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Evaluation IDs.
func (eq *EvaluationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(evaluation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EvaluationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EvaluationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EvaluationQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EvaluationQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EvaluationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EvaluationQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EvaluationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EvaluationQuery) Clone() *EvaluationQuery {
	if eq == nil {
		return nil
	}
	return &EvaluationQuery{
		config:       eq.config,
		ctx:          eq.ctx.Clone(),
		order:        append([]evaluation.OrderOption{}, eq.order...),
		inters:       append([]Interceptor{}, eq.inters...),
		predicates:   append([]predicate.Evaluation{}, eq.predicates...),
		withUser:     eq.withUser.Clone(),
		withResponse: eq.withResponse.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EvaluationQuery) WithUser(opts ...func(*UserQuery)) *EvaluationQuery {
	query := (&UserClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withUser = query
	return eq
}

// WithResponse tells the query-builder to eager-load the nodes that are connected to
// the "response" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EvaluationQuery) WithResponse(opts ...func(*ResponseQuery)) *EvaluationQuery {
	query := (&ResponseClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withResponse = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserId uuid.UUID `json:"userId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Evaluation.Query().
//		GroupBy(evaluation.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EvaluationQuery) GroupBy(field string, fields ...string) *EvaluationGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EvaluationGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = evaluation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserId uuid.UUID `json:"userId,omitempty"`
//	}
//
//	client.Evaluation.Query().
//		Select(evaluation.FieldUserId).
//		Scan(ctx, &v)
func (eq *EvaluationQuery) Select(fields ...string) *EvaluationSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EvaluationSelect{EvaluationQuery: eq}
	sbuild.label = evaluation.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EvaluationSelect configured with the given aggregations.
func (eq *EvaluationQuery) Aggregate(fns ...AggregateFunc) *EvaluationSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EvaluationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !evaluation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EvaluationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Evaluation, error) {
	var (
		nodes       = []*Evaluation{}
		_spec       = eq.querySpec()
		loadedTypes = [2]bool{
			eq.withUser != nil,
			eq.withResponse != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Evaluation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Evaluation{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withUser; query != nil {
		if err := eq.loadUser(ctx, query, nodes, nil,
			func(n *Evaluation, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withResponse; query != nil {
		if err := eq.loadResponse(ctx, query, nodes, nil,
			func(n *Evaluation, e *Response) { n.Edges.Response = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EvaluationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Evaluation, init func(*Evaluation), assign func(*Evaluation, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Evaluation)
	for i := range nodes {
		fk := nodes[i].UserId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "userId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (eq *EvaluationQuery) loadResponse(ctx context.Context, query *ResponseQuery, nodes []*Evaluation, init func(*Evaluation), assign func(*Evaluation, *Response)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Evaluation)
	for i := range nodes {
		fk := nodes[i].ResponseId
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(response.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "responseId" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (eq *EvaluationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EvaluationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(evaluation.Table, evaluation.Columns, sqlgraph.NewFieldSpec(evaluation.FieldID, field.TypeUUID))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, evaluation.FieldID)
		for i := range fields {
			if fields[i] != evaluation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if eq.withUser != nil {
			_spec.Node.AddColumnOnce(evaluation.FieldUserId)
		}
		if eq.withResponse != nil {
			_spec.Node.AddColumnOnce(evaluation.FieldResponseId)
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EvaluationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(evaluation.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = evaluation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EvaluationGroupBy is the group-by builder for Evaluation entities.
type EvaluationGroupBy struct {
	selector
	build *EvaluationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EvaluationGroupBy) Aggregate(fns ...AggregateFunc) *EvaluationGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EvaluationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EvaluationQuery, *EvaluationGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EvaluationGroupBy) sqlScan(ctx context.Context, root *EvaluationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EvaluationSelect is the builder for selecting fields of Evaluation entities.
type EvaluationSelect struct {
	*EvaluationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EvaluationSelect) Aggregate(fns ...AggregateFunc) *EvaluationSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EvaluationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EvaluationQuery, *EvaluationSelect](ctx, es.EvaluationQuery, es, es.inters, v)
}

func (es *EvaluationSelect) sqlScan(ctx context.Context, root *EvaluationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}