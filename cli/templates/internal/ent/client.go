// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/transerver/cli/templates/internal/ent/migrate"

	"github.com/transerver/cli/templates/internal/ent/greeter"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Greeter is the client for interacting with the Greeter builders.
	Greeter *GreeterClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Greeter = NewGreeterClient(c.config)
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
		ctx:     ctx,
		config:  cfg,
		Greeter: NewGreeterClient(cfg),
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
		Greeter: NewGreeterClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Greeter.
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
	c.Greeter.Use(hooks...)
}

// GreeterClient is a client for the Greeter schema.
type GreeterClient struct {
	config
}

// NewGreeterClient returns a client for the Greeter from the given config.
func NewGreeterClient(c config) *GreeterClient {
	return &GreeterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `greeter.Hooks(f(g(h())))`.
func (c *GreeterClient) Use(hooks ...Hook) {
	c.hooks.Greeter = append(c.hooks.Greeter, hooks...)
}

// Create returns a builder for creating a Greeter entity.
func (c *GreeterClient) Create() *GreeterCreate {
	mutation := newGreeterMutation(c.config, OpCreate)
	return &GreeterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Greeter entities.
func (c *GreeterClient) CreateBulk(builders ...*GreeterCreate) *GreeterCreateBulk {
	return &GreeterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Greeter.
func (c *GreeterClient) Update() *GreeterUpdate {
	mutation := newGreeterMutation(c.config, OpUpdate)
	return &GreeterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GreeterClient) UpdateOne(gr *Greeter) *GreeterUpdateOne {
	mutation := newGreeterMutation(c.config, OpUpdateOne, withGreeter(gr))
	return &GreeterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GreeterClient) UpdateOneID(id int) *GreeterUpdateOne {
	mutation := newGreeterMutation(c.config, OpUpdateOne, withGreeterID(id))
	return &GreeterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Greeter.
func (c *GreeterClient) Delete() *GreeterDelete {
	mutation := newGreeterMutation(c.config, OpDelete)
	return &GreeterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GreeterClient) DeleteOne(gr *Greeter) *GreeterDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GreeterClient) DeleteOneID(id int) *GreeterDeleteOne {
	builder := c.Delete().Where(greeter.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GreeterDeleteOne{builder}
}

// Query returns a query builder for Greeter.
func (c *GreeterClient) Query() *GreeterQuery {
	return &GreeterQuery{
		config: c.config,
	}
}

// Get returns a Greeter entity by its id.
func (c *GreeterClient) Get(ctx context.Context, id int) (*Greeter, error) {
	return c.Query().Where(greeter.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GreeterClient) GetX(ctx context.Context, id int) *Greeter {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GreeterClient) Hooks() []Hook {
	return c.hooks.Greeter
}
