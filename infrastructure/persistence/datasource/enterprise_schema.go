package datasource

import (
	"context"
	"sync"
)

var enterpriseSchema *Schema
var once3 sync.Once

func EnterpriseSchema() *Schema {
	var err error = nil
	once3.Do(func() {
		ctx := context.Background()
		enterpriseSchema, err = NewPGSqlSchema(ctx, WithSchema("enterprise"))
		if err != nil {
			panic(err)
		}
		err = enterpriseSchema.Connect(ctx)
		if err != nil {
			panic(err)
		}
	})
	return enterpriseSchema
}
