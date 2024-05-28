package datasource

import (
	"context"
	"sync"
)

var eventSchema *Schema
var once2 sync.Once

func EventSchema() *Schema {
	var err error = nil
	once2.Do(func() {
		ctx := context.Background()
		eventSchema, err = NewPGSqlSchema(ctx, WithSchema("event"))
		if err != nil {
			panic(err)
		}
		err = eventSchema.Connect(ctx)
		if err != nil {
			panic(err)
		}
	})
	return eventSchema
}
