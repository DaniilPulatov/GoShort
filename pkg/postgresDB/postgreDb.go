package postgresDB

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Pool interface {
	Ping(ctx context.Context) error
	Close()

	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type conn struct {
	*pgxpool.Pool
}

func NewPostgresDB(dsn string) (Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Println("Error parsing DSN:", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Println("Error connecting to database:", err)
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Println("Failed to ping pool:", err)
		return nil, fmt.Errorf("failed to ping pool: %w", err)
	}
	log.Println("Database connection established successfully")
	return &conn{
		Pool: pool,
	}, nil
}

func (c *conn) Ping(ctx context.Context) error {
	if err := c.Pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping pool: %w", err)
	}
	return nil
}

func (c *conn) Close() {
	c.Pool.Close()
}

func (c *conn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return c.Pool.QueryRow(ctx, sql, args...)
}

func (c *conn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	rows, err := c.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	return rows, nil
}

func (c *conn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	tag, err := c.Pool.Exec(ctx, sql, arguments...)
	if err != nil {
		return tag, fmt.Errorf("failed to exec: %w", err)
	}
	return tag, nil
}
