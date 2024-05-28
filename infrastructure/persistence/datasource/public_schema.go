package datasource

import (
	"context"
	"sync"
)

var publicSchema *Schema
var once5 sync.Once

func PublicSchema() *Schema {
	var err error = nil
	once5.Do(func() {
		ctx := context.Background()
		publicSchema, err = NewPGSqlSchema(ctx)
		if err != nil {
			panic(err)
		}
		err = publicSchema.Connect(ctx)
		if err != nil {
			panic(err)
		}
	})
	return publicSchema
}
