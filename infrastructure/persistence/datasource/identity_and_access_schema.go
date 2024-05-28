package datasource

import (
	"context"
	"sync"
)

var identityandaccess *Schema
var once sync.Once

func IdentityAndAccessSchema() *Schema {
	var err error = nil
	once.Do(func() {
		ctx := context.Background()
		identityandaccess, err = NewPGSqlSchema(ctx, WithSchema("identityandaccess"))
		if err != nil {
			panic(err)
		}
		err := identityandaccess.Connect(ctx)
		if err != nil {
			panic(err)
		}
	})
	return identityandaccess
}
