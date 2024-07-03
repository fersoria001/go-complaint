package datasource

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Config() *pgxpool.Config {
	var defaultMaxConns = int32(100)
	var defaultMinConns = int32(0)
	var defaultMaxConnLifetime = time.Hour
	var defaultMaxConnIdleTime = time.Minute * 30
	var defaultHealthCheckPeriod = time.Minute
	var defaultConnectTimeout = time.Second * 5
	var DATABASE_URL string = os.Getenv("DATABASE_URL")
	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout
	if err != nil {
		return nil
	}
	return dbConfig
}

type Schema interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	Close()
}

type PGSQLSchema struct {
	pool *pgxpool.Pool
}

func NewPGSqlSchema(ctx context.Context) (*PGSQLSchema, error) {
	p := &PGSQLSchema{}
	pool, err := pgxpool.NewWithConfig(ctx, Config())
	if err != nil {
		return p, err
	}
	p.pool = pool
	return p, nil
}

func (p PGSQLSchema) Close() {
	p.pool.Close()
}
func (p PGSQLSchema) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	return p.pool.Acquire(ctx)

}

func (p *PGSQLSchema) ConnectionString() string {
	return os.Getenv("DATABASE_URL")
}
