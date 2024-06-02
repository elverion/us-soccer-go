// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"us-soccer-go-test/internal/ent/migrate"

	"us-soccer-go-test/internal/ent/stadium"
	"us-soccer-go-test/internal/ent/weather"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Stadium is the client for interacting with the Stadium builders.
	Stadium *StadiumClient
	// Weather is the client for interacting with the Weather builders.
	Weather *WeatherClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Stadium = NewStadiumClient(c.config)
	c.Weather = NewWeatherClient(c.config)
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

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Stadium: NewStadiumClient(cfg),
		Weather: NewWeatherClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Stadium: NewStadiumClient(cfg),
		Weather: NewWeatherClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Stadium.
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
	c.Stadium.Use(hooks...)
	c.Weather.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Stadium.Intercept(interceptors...)
	c.Weather.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *StadiumMutation:
		return c.Stadium.mutate(ctx, m)
	case *WeatherMutation:
		return c.Weather.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// StadiumClient is a client for the Stadium schema.
type StadiumClient struct {
	config
}

// NewStadiumClient returns a client for the Stadium from the given config.
func NewStadiumClient(c config) *StadiumClient {
	return &StadiumClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `stadium.Hooks(f(g(h())))`.
func (c *StadiumClient) Use(hooks ...Hook) {
	c.hooks.Stadium = append(c.hooks.Stadium, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `stadium.Intercept(f(g(h())))`.
func (c *StadiumClient) Intercept(interceptors ...Interceptor) {
	c.inters.Stadium = append(c.inters.Stadium, interceptors...)
}

// Create returns a builder for creating a Stadium entity.
func (c *StadiumClient) Create() *StadiumCreate {
	mutation := newStadiumMutation(c.config, OpCreate)
	return &StadiumCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Stadium entities.
func (c *StadiumClient) CreateBulk(builders ...*StadiumCreate) *StadiumCreateBulk {
	return &StadiumCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StadiumClient) MapCreateBulk(slice any, setFunc func(*StadiumCreate, int)) *StadiumCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StadiumCreateBulk{err: fmt.Errorf("calling to StadiumClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StadiumCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StadiumCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Stadium.
func (c *StadiumClient) Update() *StadiumUpdate {
	mutation := newStadiumMutation(c.config, OpUpdate)
	return &StadiumUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StadiumClient) UpdateOne(s *Stadium) *StadiumUpdateOne {
	mutation := newStadiumMutation(c.config, OpUpdateOne, withStadium(s))
	return &StadiumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StadiumClient) UpdateOneID(id uuid.UUID) *StadiumUpdateOne {
	mutation := newStadiumMutation(c.config, OpUpdateOne, withStadiumID(id))
	return &StadiumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Stadium.
func (c *StadiumClient) Delete() *StadiumDelete {
	mutation := newStadiumMutation(c.config, OpDelete)
	return &StadiumDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StadiumClient) DeleteOne(s *Stadium) *StadiumDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StadiumClient) DeleteOneID(id uuid.UUID) *StadiumDeleteOne {
	builder := c.Delete().Where(stadium.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StadiumDeleteOne{builder}
}

// Query returns a query builder for Stadium.
func (c *StadiumClient) Query() *StadiumQuery {
	return &StadiumQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStadium},
		inters: c.Interceptors(),
	}
}

// Get returns a Stadium entity by its id.
func (c *StadiumClient) Get(ctx context.Context, id uuid.UUID) (*Stadium, error) {
	return c.Query().Where(stadium.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StadiumClient) GetX(ctx context.Context, id uuid.UUID) *Stadium {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWeather queries the weather edge of a Stadium.
func (c *StadiumClient) QueryWeather(s *Stadium) *WeatherQuery {
	query := (&WeatherClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(stadium.Table, stadium.FieldID, id),
			sqlgraph.To(weather.Table, weather.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, stadium.WeatherTable, stadium.WeatherColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StadiumClient) Hooks() []Hook {
	hooks := c.hooks.Stadium
	return append(hooks[:len(hooks):len(hooks)], stadium.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *StadiumClient) Interceptors() []Interceptor {
	return c.inters.Stadium
}

func (c *StadiumClient) mutate(ctx context.Context, m *StadiumMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StadiumCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StadiumUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StadiumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StadiumDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Stadium mutation op: %q", m.Op())
	}
}

// WeatherClient is a client for the Weather schema.
type WeatherClient struct {
	config
}

// NewWeatherClient returns a client for the Weather from the given config.
func NewWeatherClient(c config) *WeatherClient {
	return &WeatherClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `weather.Hooks(f(g(h())))`.
func (c *WeatherClient) Use(hooks ...Hook) {
	c.hooks.Weather = append(c.hooks.Weather, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `weather.Intercept(f(g(h())))`.
func (c *WeatherClient) Intercept(interceptors ...Interceptor) {
	c.inters.Weather = append(c.inters.Weather, interceptors...)
}

// Create returns a builder for creating a Weather entity.
func (c *WeatherClient) Create() *WeatherCreate {
	mutation := newWeatherMutation(c.config, OpCreate)
	return &WeatherCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Weather entities.
func (c *WeatherClient) CreateBulk(builders ...*WeatherCreate) *WeatherCreateBulk {
	return &WeatherCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *WeatherClient) MapCreateBulk(slice any, setFunc func(*WeatherCreate, int)) *WeatherCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &WeatherCreateBulk{err: fmt.Errorf("calling to WeatherClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*WeatherCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &WeatherCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Weather.
func (c *WeatherClient) Update() *WeatherUpdate {
	mutation := newWeatherMutation(c.config, OpUpdate)
	return &WeatherUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WeatherClient) UpdateOne(w *Weather) *WeatherUpdateOne {
	mutation := newWeatherMutation(c.config, OpUpdateOne, withWeather(w))
	return &WeatherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WeatherClient) UpdateOneID(id uuid.UUID) *WeatherUpdateOne {
	mutation := newWeatherMutation(c.config, OpUpdateOne, withWeatherID(id))
	return &WeatherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Weather.
func (c *WeatherClient) Delete() *WeatherDelete {
	mutation := newWeatherMutation(c.config, OpDelete)
	return &WeatherDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WeatherClient) DeleteOne(w *Weather) *WeatherDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WeatherClient) DeleteOneID(id uuid.UUID) *WeatherDeleteOne {
	builder := c.Delete().Where(weather.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WeatherDeleteOne{builder}
}

// Query returns a query builder for Weather.
func (c *WeatherClient) Query() *WeatherQuery {
	return &WeatherQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeWeather},
		inters: c.Interceptors(),
	}
}

// Get returns a Weather entity by its id.
func (c *WeatherClient) Get(ctx context.Context, id uuid.UUID) (*Weather, error) {
	return c.Query().Where(weather.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WeatherClient) GetX(ctx context.Context, id uuid.UUID) *Weather {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStadium queries the stadium edge of a Weather.
func (c *WeatherClient) QueryStadium(w *Weather) *StadiumQuery {
	query := (&StadiumClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(weather.Table, weather.FieldID, id),
			sqlgraph.To(stadium.Table, stadium.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, weather.StadiumTable, weather.StadiumColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WeatherClient) Hooks() []Hook {
	hooks := c.hooks.Weather
	return append(hooks[:len(hooks):len(hooks)], weather.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *WeatherClient) Interceptors() []Interceptor {
	return c.inters.Weather
}

func (c *WeatherClient) mutate(ctx context.Context, m *WeatherMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&WeatherCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&WeatherUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&WeatherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&WeatherDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Weather mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Stadium, Weather []ent.Hook
	}
	inters struct {
		Stadium, Weather []ent.Interceptor
	}
)
