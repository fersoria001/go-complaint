package datasource

import (
	"context"
	"fmt"
	"go-complaint/erros"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Schema struct {
	host       string
	port       int
	user       string
	dbname     string
	schemaName string
	maxConn    int
	Pool       *pgxpool.Pool
}

func NewPGSqlSchema(ctx context.Context, options ...OptionsPGSqlSchemaFunc) (*Schema, error) {
	host := "localhost"
	port := 5432
	user := "postgres"
	dbname := "postgres"
	schemaName := "public"
	maxConn := 10
	p := &Schema{
		host:       host,
		port:       port,
		user:       user,
		dbname:     dbname,
		schemaName: schemaName,
		maxConn:    maxConn,
	}
	for _, option := range options {
		err := option(p)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

type OptionsPGSqlSchemaFunc func(p *Schema) error

func (p *Schema) Connect(ctx context.Context) error {
	var err error
	p.Pool, err = pgxpool.New(ctx, p.ConnectionString())
	if err != nil {
		return err
	}
	return nil
}

func (p *Schema) Close() {
	p.Pool.Close()
}

func (p *Schema) ConnectionString() string {
	password := "sfdkwtf"
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?pool_max_conns=%d&search_path=%s&connect_timeout=5", p.user, password, p.host, p.port, p.dbname, p.maxConn, p.schemaName)
}

func WithHost(host string) OptionsPGSqlSchemaFunc {
	return func(p *Schema) error {
		p.host = host
		return nil
	}
}

func WithPort(port int) OptionsPGSqlSchemaFunc {
	return func(p *Schema) error {
		if port < 0 || port > 65535 {
			return &erros.InvalidPortError{Port: port}
		}
		p.port = port
		return nil
	}
}

func WithUser(user string) OptionsPGSqlSchemaFunc {
	return func(p *Schema) error {
		p.user = user
		return nil
	}
}

func WithDbname(dbname string) OptionsPGSqlSchemaFunc {
	return func(p *Schema) error {
		p.dbname = dbname
		return nil
	}
}

func WithSchema(schemaName string) OptionsPGSqlSchemaFunc {
	return func(p *Schema) error {
		p.schemaName = schemaName
		return nil
	}
}

func WithMaxConn(maxConn int) OptionsPGSqlSchemaFunc {
	return func(p *Schema) error {
		p.maxConn = maxConn
		return nil
	}
}

func (p *Schema) SchemaName() string {
	return p.schemaName
}
