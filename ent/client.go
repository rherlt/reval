// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/rherlt/reval/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rherlt/reval/ent/evaluation"
	"github.com/rherlt/reval/ent/request"
	"github.com/rherlt/reval/ent/response"
	"github.com/rherlt/reval/ent/scenario"
	"github.com/rherlt/reval/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Evaluation is the client for interacting with the Evaluation builders.
	Evaluation *EvaluationClient
	// Request is the client for interacting with the Request builders.
	Request *RequestClient
	// Response is the client for interacting with the Response builders.
	Response *ResponseClient
	// Scenario is the client for interacting with the Scenario builders.
	Scenario *ScenarioClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Evaluation = NewEvaluationClient(c.config)
	c.Request = NewRequestClient(c.config)
	c.Response = NewResponseClient(c.config)
	c.Scenario = NewScenarioClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Evaluation: NewEvaluationClient(cfg),
		Request:    NewRequestClient(cfg),
		Response:   NewResponseClient(cfg),
		Scenario:   NewScenarioClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Evaluation: NewEvaluationClient(cfg),
		Request:    NewRequestClient(cfg),
		Response:   NewResponseClient(cfg),
		Scenario:   NewScenarioClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Evaluation.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Evaluation.Use(hooks...)
	c.Request.Use(hooks...)
	c.Response.Use(hooks...)
	c.Scenario.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Evaluation.Intercept(interceptors...)
	c.Request.Intercept(interceptors...)
	c.Response.Intercept(interceptors...)
	c.Scenario.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *EvaluationMutation:
		return c.Evaluation.mutate(ctx, m)
	case *RequestMutation:
		return c.Request.mutate(ctx, m)
	case *ResponseMutation:
		return c.Response.mutate(ctx, m)
	case *ScenarioMutation:
		return c.Scenario.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// EvaluationClient is a client for the Evaluation schema.
type EvaluationClient struct {
	config
}

// NewEvaluationClient returns a client for the Evaluation from the given config.
func NewEvaluationClient(c config) *EvaluationClient {
	return &EvaluationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `evaluation.Hooks(f(g(h())))`.
func (c *EvaluationClient) Use(hooks ...Hook) {
	c.hooks.Evaluation = append(c.hooks.Evaluation, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `evaluation.Intercept(f(g(h())))`.
func (c *EvaluationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Evaluation = append(c.inters.Evaluation, interceptors...)
}

// Create returns a builder for creating a Evaluation entity.
func (c *EvaluationClient) Create() *EvaluationCreate {
	mutation := newEvaluationMutation(c.config, OpCreate)
	return &EvaluationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Evaluation entities.
func (c *EvaluationClient) CreateBulk(builders ...*EvaluationCreate) *EvaluationCreateBulk {
	return &EvaluationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Evaluation.
func (c *EvaluationClient) Update() *EvaluationUpdate {
	mutation := newEvaluationMutation(c.config, OpUpdate)
	return &EvaluationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EvaluationClient) UpdateOne(e *Evaluation) *EvaluationUpdateOne {
	mutation := newEvaluationMutation(c.config, OpUpdateOne, withEvaluation(e))
	return &EvaluationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EvaluationClient) UpdateOneID(id uuid.UUID) *EvaluationUpdateOne {
	mutation := newEvaluationMutation(c.config, OpUpdateOne, withEvaluationID(id))
	return &EvaluationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Evaluation.
func (c *EvaluationClient) Delete() *EvaluationDelete {
	mutation := newEvaluationMutation(c.config, OpDelete)
	return &EvaluationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EvaluationClient) DeleteOne(e *Evaluation) *EvaluationDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EvaluationClient) DeleteOneID(id uuid.UUID) *EvaluationDeleteOne {
	builder := c.Delete().Where(evaluation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EvaluationDeleteOne{builder}
}

// Query returns a query builder for Evaluation.
func (c *EvaluationClient) Query() *EvaluationQuery {
	return &EvaluationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEvaluation},
		inters: c.Interceptors(),
	}
}

// Get returns a Evaluation entity by its id.
func (c *EvaluationClient) Get(ctx context.Context, id uuid.UUID) (*Evaluation, error) {
	return c.Query().Where(evaluation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EvaluationClient) GetX(ctx context.Context, id uuid.UUID) *Evaluation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Evaluation.
func (c *EvaluationClient) QueryUser(e *Evaluation) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(evaluation.Table, evaluation.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, evaluation.UserTable, evaluation.UserColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryResponse queries the response edge of a Evaluation.
func (c *EvaluationClient) QueryResponse(e *Evaluation) *ResponseQuery {
	query := (&ResponseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(evaluation.Table, evaluation.FieldID, id),
			sqlgraph.To(response.Table, response.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, evaluation.ResponseTable, evaluation.ResponseColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EvaluationClient) Hooks() []Hook {
	return c.hooks.Evaluation
}

// Interceptors returns the client interceptors.
func (c *EvaluationClient) Interceptors() []Interceptor {
	return c.inters.Evaluation
}

func (c *EvaluationClient) mutate(ctx context.Context, m *EvaluationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EvaluationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EvaluationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EvaluationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EvaluationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Evaluation mutation op: %q", m.Op())
	}
}

// RequestClient is a client for the Request schema.
type RequestClient struct {
	config
}

// NewRequestClient returns a client for the Request from the given config.
func NewRequestClient(c config) *RequestClient {
	return &RequestClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `request.Hooks(f(g(h())))`.
func (c *RequestClient) Use(hooks ...Hook) {
	c.hooks.Request = append(c.hooks.Request, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `request.Intercept(f(g(h())))`.
func (c *RequestClient) Intercept(interceptors ...Interceptor) {
	c.inters.Request = append(c.inters.Request, interceptors...)
}

// Create returns a builder for creating a Request entity.
func (c *RequestClient) Create() *RequestCreate {
	mutation := newRequestMutation(c.config, OpCreate)
	return &RequestCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Request entities.
func (c *RequestClient) CreateBulk(builders ...*RequestCreate) *RequestCreateBulk {
	return &RequestCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Request.
func (c *RequestClient) Update() *RequestUpdate {
	mutation := newRequestMutation(c.config, OpUpdate)
	return &RequestUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RequestClient) UpdateOne(r *Request) *RequestUpdateOne {
	mutation := newRequestMutation(c.config, OpUpdateOne, withRequest(r))
	return &RequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RequestClient) UpdateOneID(id uuid.UUID) *RequestUpdateOne {
	mutation := newRequestMutation(c.config, OpUpdateOne, withRequestID(id))
	return &RequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Request.
func (c *RequestClient) Delete() *RequestDelete {
	mutation := newRequestMutation(c.config, OpDelete)
	return &RequestDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RequestClient) DeleteOne(r *Request) *RequestDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RequestClient) DeleteOneID(id uuid.UUID) *RequestDeleteOne {
	builder := c.Delete().Where(request.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RequestDeleteOne{builder}
}

// Query returns a query builder for Request.
func (c *RequestClient) Query() *RequestQuery {
	return &RequestQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRequest},
		inters: c.Interceptors(),
	}
}

// Get returns a Request entity by its id.
func (c *RequestClient) Get(ctx context.Context, id uuid.UUID) (*Request, error) {
	return c.Query().Where(request.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RequestClient) GetX(ctx context.Context, id uuid.UUID) *Request {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryResponses queries the responses edge of a Request.
func (c *RequestClient) QueryResponses(r *Request) *ResponseQuery {
	query := (&ResponseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(request.Table, request.FieldID, id),
			sqlgraph.To(response.Table, response.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, request.ResponsesTable, request.ResponsesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RequestClient) Hooks() []Hook {
	return c.hooks.Request
}

// Interceptors returns the client interceptors.
func (c *RequestClient) Interceptors() []Interceptor {
	return c.inters.Request
}

func (c *RequestClient) mutate(ctx context.Context, m *RequestMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RequestCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RequestUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RequestDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Request mutation op: %q", m.Op())
	}
}

// ResponseClient is a client for the Response schema.
type ResponseClient struct {
	config
}

// NewResponseClient returns a client for the Response from the given config.
func NewResponseClient(c config) *ResponseClient {
	return &ResponseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `response.Hooks(f(g(h())))`.
func (c *ResponseClient) Use(hooks ...Hook) {
	c.hooks.Response = append(c.hooks.Response, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `response.Intercept(f(g(h())))`.
func (c *ResponseClient) Intercept(interceptors ...Interceptor) {
	c.inters.Response = append(c.inters.Response, interceptors...)
}

// Create returns a builder for creating a Response entity.
func (c *ResponseClient) Create() *ResponseCreate {
	mutation := newResponseMutation(c.config, OpCreate)
	return &ResponseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Response entities.
func (c *ResponseClient) CreateBulk(builders ...*ResponseCreate) *ResponseCreateBulk {
	return &ResponseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Response.
func (c *ResponseClient) Update() *ResponseUpdate {
	mutation := newResponseMutation(c.config, OpUpdate)
	return &ResponseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ResponseClient) UpdateOne(r *Response) *ResponseUpdateOne {
	mutation := newResponseMutation(c.config, OpUpdateOne, withResponse(r))
	return &ResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ResponseClient) UpdateOneID(id uuid.UUID) *ResponseUpdateOne {
	mutation := newResponseMutation(c.config, OpUpdateOne, withResponseID(id))
	return &ResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Response.
func (c *ResponseClient) Delete() *ResponseDelete {
	mutation := newResponseMutation(c.config, OpDelete)
	return &ResponseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ResponseClient) DeleteOne(r *Response) *ResponseDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ResponseClient) DeleteOneID(id uuid.UUID) *ResponseDeleteOne {
	builder := c.Delete().Where(response.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ResponseDeleteOne{builder}
}

// Query returns a query builder for Response.
func (c *ResponseClient) Query() *ResponseQuery {
	return &ResponseQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeResponse},
		inters: c.Interceptors(),
	}
}

// Get returns a Response entity by its id.
func (c *ResponseClient) Get(ctx context.Context, id uuid.UUID) (*Response, error) {
	return c.Query().Where(response.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ResponseClient) GetX(ctx context.Context, id uuid.UUID) *Response {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRequest queries the request edge of a Response.
func (c *ResponseClient) QueryRequest(r *Response) *RequestQuery {
	query := (&RequestClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(response.Table, response.FieldID, id),
			sqlgraph.To(request.Table, request.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, response.RequestTable, response.RequestColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryScenario queries the scenario edge of a Response.
func (c *ResponseClient) QueryScenario(r *Response) *ScenarioQuery {
	query := (&ScenarioClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(response.Table, response.FieldID, id),
			sqlgraph.To(scenario.Table, scenario.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, response.ScenarioTable, response.ScenarioColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEvaluations queries the evaluations edge of a Response.
func (c *ResponseClient) QueryEvaluations(r *Response) *EvaluationQuery {
	query := (&EvaluationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(response.Table, response.FieldID, id),
			sqlgraph.To(evaluation.Table, evaluation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, response.EvaluationsTable, response.EvaluationsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ResponseClient) Hooks() []Hook {
	return c.hooks.Response
}

// Interceptors returns the client interceptors.
func (c *ResponseClient) Interceptors() []Interceptor {
	return c.inters.Response
}

func (c *ResponseClient) mutate(ctx context.Context, m *ResponseMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ResponseCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ResponseUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ResponseDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Response mutation op: %q", m.Op())
	}
}

// ScenarioClient is a client for the Scenario schema.
type ScenarioClient struct {
	config
}

// NewScenarioClient returns a client for the Scenario from the given config.
func NewScenarioClient(c config) *ScenarioClient {
	return &ScenarioClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `scenario.Hooks(f(g(h())))`.
func (c *ScenarioClient) Use(hooks ...Hook) {
	c.hooks.Scenario = append(c.hooks.Scenario, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `scenario.Intercept(f(g(h())))`.
func (c *ScenarioClient) Intercept(interceptors ...Interceptor) {
	c.inters.Scenario = append(c.inters.Scenario, interceptors...)
}

// Create returns a builder for creating a Scenario entity.
func (c *ScenarioClient) Create() *ScenarioCreate {
	mutation := newScenarioMutation(c.config, OpCreate)
	return &ScenarioCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Scenario entities.
func (c *ScenarioClient) CreateBulk(builders ...*ScenarioCreate) *ScenarioCreateBulk {
	return &ScenarioCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Scenario.
func (c *ScenarioClient) Update() *ScenarioUpdate {
	mutation := newScenarioMutation(c.config, OpUpdate)
	return &ScenarioUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ScenarioClient) UpdateOne(s *Scenario) *ScenarioUpdateOne {
	mutation := newScenarioMutation(c.config, OpUpdateOne, withScenario(s))
	return &ScenarioUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ScenarioClient) UpdateOneID(id uuid.UUID) *ScenarioUpdateOne {
	mutation := newScenarioMutation(c.config, OpUpdateOne, withScenarioID(id))
	return &ScenarioUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Scenario.
func (c *ScenarioClient) Delete() *ScenarioDelete {
	mutation := newScenarioMutation(c.config, OpDelete)
	return &ScenarioDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ScenarioClient) DeleteOne(s *Scenario) *ScenarioDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ScenarioClient) DeleteOneID(id uuid.UUID) *ScenarioDeleteOne {
	builder := c.Delete().Where(scenario.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ScenarioDeleteOne{builder}
}

// Query returns a query builder for Scenario.
func (c *ScenarioClient) Query() *ScenarioQuery {
	return &ScenarioQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeScenario},
		inters: c.Interceptors(),
	}
}

// Get returns a Scenario entity by its id.
func (c *ScenarioClient) Get(ctx context.Context, id uuid.UUID) (*Scenario, error) {
	return c.Query().Where(scenario.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ScenarioClient) GetX(ctx context.Context, id uuid.UUID) *Scenario {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryResponses queries the responses edge of a Scenario.
func (c *ScenarioClient) QueryResponses(s *Scenario) *ResponseQuery {
	query := (&ResponseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(scenario.Table, scenario.FieldID, id),
			sqlgraph.To(response.Table, response.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scenario.ResponsesTable, scenario.ResponsesColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ScenarioClient) Hooks() []Hook {
	return c.hooks.Scenario
}

// Interceptors returns the client interceptors.
func (c *ScenarioClient) Interceptors() []Interceptor {
	return c.inters.Scenario
}

func (c *ScenarioClient) mutate(ctx context.Context, m *ScenarioMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ScenarioCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ScenarioUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ScenarioUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ScenarioDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Scenario mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEvaluations queries the evaluations edge of a User.
func (c *UserClient) QueryEvaluations(u *User) *EvaluationQuery {
	query := (&EvaluationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(evaluation.Table, evaluation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.EvaluationsTable, user.EvaluationsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Evaluation, Request, Response, Scenario, User []ent.Hook
	}
	inters struct {
		Evaluation, Request, Response, Scenario, User []ent.Interceptor
	}
)
