package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type AddressRepository struct {
	schema datasource.Schema
}

func NewAddressRepository(schema datasource.Schema) AddressRepository {
	return AddressRepository{schema: schema}
}

func (ar AddressRepository) Update(
	ctx context.Context,
	address common.Address,
) error {
	conn, err := ar.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	updateQuery := string(`
		UPDATE address
		SET
			country_id = $2,
			country_state_id = $3,
			state_city_id = $4
		WHERE id = $1
	`)
	var (
		id           uuid.UUID = address.ID()
		countryID    int       = address.Country().ID()
		countryState int       = address.CountryState().ID()
		stateCity    int       = address.City().ID()
	)
	_, err = conn.Exec(
		ctx,
		updateQuery,
		id,
		countryID,
		countryState,
		stateCity,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}
func (ar AddressRepository) Remove(
	ctx context.Context,
	id uuid.UUID,
) error {
	conn, err := ar.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "DELETE FROM ADDRESS WHERE ID=$1", &id)
	if err != nil {
		return err
	}
	return nil
}
func (ar AddressRepository) Save(
	ctx context.Context,
	address common.Address,
) error {
	conn, err := ar.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertQuery := string(`
		INSERT INTO address (
			id,
			country_id,
			country_state_id,
			state_city_id
		) VALUES ($1, $2, $3, $4)
		 `)
	var (
		id           uuid.UUID = address.ID()
		countryID    int       = address.Country().ID()
		countryState int       = address.CountryState().ID()
		stateCity    int       = address.City().ID()
	)
	_, err = conn.Exec(
		ctx,
		insertQuery,
		id,
		countryID,
		countryState,
		stateCity,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (ar AddressRepository) Get(
	ctx context.Context,
	id uuid.UUID,
) (common.Address, error) {
	conn, err := ar.schema.Acquire(ctx)
	if err != nil {
		return common.Address{}, err
	}
	selectQuery := string(`
		SELECT
		id,
		country_id,
		country_state_id,
		state_city_id
		FROM address
		WHERE id = $1`,
	)
	row := conn.QueryRow(
		ctx,
		selectQuery,
		id,
	)
	address, err := ar.load(ctx, row)
	if err != nil {
		return common.Address{}, err
	}
	defer conn.Release()
	return address, nil
}

func (ar AddressRepository) load(
	ctx context.Context,
	row pgx.Row,
) (common.Address, error) {
	mapper := MapperRegistryInstance().Get("Country")
	if mapper == nil {
		return common.Address{}, ErrMapperNotRegistered
	}
	countryRepository, ok := mapper.(CountryRepository)
	if !ok {
		return common.Address{}, ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("CountryState")
	if mapper == nil {
		return common.Address{}, ErrMapperNotRegistered
	}
	countryStatesRepository, ok := mapper.(CountryStateRepository)
	if !ok {
		return common.Address{}, ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("City")
	if mapper == nil {
		return common.Address{}, ErrMapperNotRegistered
	}
	stateCitiesRepository, ok := mapper.(StateCitiesRepository)
	if !ok {
		return common.Address{}, ErrWrongTypeAssertion
	}
	var (
		id             uuid.UUID
		countryID      int
		countryStateID int
		stateCityID    int
	)
	err := row.Scan(
		&id,
		&countryID,
		&countryStateID,
		&stateCityID,
	)
	if err != nil {
		return common.Address{}, err
	}
	country, err := countryRepository.Get(ctx, countryID)
	if err != nil {
		return common.Address{}, err
	}
	countryState, err := countryStatesRepository.Get(ctx, countryStateID)
	if err != nil {
		return common.Address{}, err
	}
	stateCity, err := stateCitiesRepository.Get(ctx, stateCityID)
	if err != nil {
		return common.Address{}, err
	}
	return common.NewAddress(
		id,
		country,
		countryState,
		stateCity,
	), nil
}
