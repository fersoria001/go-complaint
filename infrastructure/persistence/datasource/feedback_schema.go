package datasource

import (
	"context"
	"sync"
)

var feedbackSchema *Schema
var once6 sync.Once

func FeedbackSchema() *Schema {
	var err error = nil
	once6.Do(func() {
		ctx := context.Background()
		feedbackSchema, err = NewPGSqlSchema(ctx, WithSchema("feedback"))
		if err != nil {
			panic(err)
		}
		err = feedbackSchema.Connect(ctx)
		if err != nil {
			panic(err)
		}
	})
	return feedbackSchema
}
