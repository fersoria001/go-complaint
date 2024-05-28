package repositories

import (
	"context"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// Package repositories
// Collection oriented repository interface
// Specific to PostgresSQL Pgx driver
// here interface{} is a domain object
type Repository interface {
	GetAll(ctx context.Context) (mapset.Set[interface{}], error)
	Get(ctx context.Context, id uuid.UUID) (interface{}, error)
	Save(ctx context.Context, entity interface{}) error
	SaveAll(ctx context.Context, entities mapset.Set[interface{}]) error
	Remove(ctx context.Context, entity interface{}) error
	RemoveAll(ctx context.Context, entities mapset.Set[interface{}]) error
}
