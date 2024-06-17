package datasource

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Schema interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	Close()
}

type PGSQLSchema struct {
	host       string
	port       int
	user       string
	dbname     string
	schemaName string
	maxConn    int
	pool       *pgxpool.Pool
}

func NewPGSqlSchema(ctx context.Context) (*PGSQLSchema, error) {
	host := "localhost"
	port := 5432
	user := "postgres"
	dbname := "postgres"
	schemaName := "public"
	maxConn := 100
	p := &PGSQLSchema{
		host:       host,
		port:       port,
		user:       user,
		dbname:     dbname,
		schemaName: schemaName,
		maxConn:    maxConn,
	}
	pool, err := pgxpool.New(ctx, p.ConnectionString())
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
	password := "sfdkwtf"
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?pool_max_conns=%d&search_path=%s&connect_timeout=5",
		p.user, password, p.host, p.port, p.dbname, p.maxConn, p.schemaName)
}
