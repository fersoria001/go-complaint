package datasource

import (
	"context"
	"sync"
)

var publicSchema Schema
var schemaOnce sync.Once

func PublicSchema() Schema {
	var err error = nil
	schemaOnce.Do(func() {
		ctx := context.Background()
		publicSchema, err = NewPGSqlSchema(ctx)
		if err != nil {
			panic(err)
		}
	})
	return publicSchema
}
