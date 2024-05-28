package datasource

import (
	"context"
	"sync"
)

var complaintSchema *Schema
var once4 sync.Once

func ComplaintSchema() *Schema {
	var err error = nil
	once4.Do(func() {
		ctx := context.Background()
		complaintSchema, err = NewPGSqlSchema(ctx, WithSchema("complaint"))
		if err != nil {
			panic(err)
		}
		err = complaintSchema.Connect(ctx)
		if err != nil {
			panic(err)
		}
	})
	return complaintSchema
}
