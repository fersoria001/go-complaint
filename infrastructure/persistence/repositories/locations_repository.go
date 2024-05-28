package repositories

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/datasource"
)

type LocationsRepository struct {
	schema *datasource.Schema
}

func NewLocationsRepository(schema *datasource.Schema) *LocationsRepository {
	return &LocationsRepository{schema: schema}
}

func (r *LocationsRepository) FindAllCountries(ctx context.Context) ([]dto.Country, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var countries []dto.Country
	var country string
	var countryID int

	rows, err := conn.Query(ctx, "SELECT id,name FROM countries")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&countryID, &country)
		if err != nil {
			return nil, err
		}
		countries = append(countries, dto.Country{ID: countryID, Name: country})
	}
	return countries, nil
}
func (r *LocationsRepository) FindCountryByID(ctx context.Context, id int) (dto.Country, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return dto.Country{}, err
	}
	defer conn.Release()
	var country string
	var countryID int
	row := conn.QueryRow(ctx, "SELECT id,name FROM countries WHERE id = $1", id)
	err = row.Scan(&countryID, &country)
	if err != nil {
		return dto.Country{}, err
	}
	return dto.Country{ID: countryID, Name: country}, nil
}
func (r *LocationsRepository) FindCountyByID(ctx context.Context, id int) (dto.County, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return dto.County{}, err
	}
	defer conn.Release()
	var state string
	var stateID int
	row := conn.QueryRow(ctx, "SELECT id,name FROM states WHERE id = $1", id)
	err = row.Scan(&stateID, &state)
	if err != nil {
		return dto.County{}, err
	}
	return dto.County{ID: stateID, Name: state}, nil
}

func (r *LocationsRepository) FindCityByID(ctx context.Context, id int) (dto.City, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return dto.City{}, err
	}
	defer conn.Release()
	var city string
	var cityID int
	row := conn.QueryRow(ctx, "SELECT id,name FROM cities WHERE id = $1", id)
	err = row.Scan(&cityID, &city)
	if err != nil {
		return dto.City{}, err
	}
	return dto.City{ID: cityID, Name: city}, nil
}

func (r *LocationsRepository) FindCountyByCountryID(ctx context.Context, country int) ([]dto.County, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var counties []dto.County
	var county string
	var countyID int
	rows, err := conn.Query(ctx, "SELECT id,name FROM states WHERE country_id = $1", country)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&countyID, &county)
		if err != nil {
			return nil, err
		}
		counties = append(counties, dto.County{ID: countyID, Name: county})
	}
	return counties, nil
}

func (r *LocationsRepository) FindCityByCountyID(ctx context.Context, county int) ([]dto.City, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var cities []dto.City
	var city string
	var cityID int
	rows, err := conn.Query(ctx, "SELECT id,name FROM cities WHERE state_id = $1", county)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&cityID, &city)
		if err != nil {
			return nil, err
		}
		cities = append(cities, dto.City{ID: cityID, Name: city})
	}
	return cities, nil
}

func (r *LocationsRepository) FindPhoneCodeByCountryID(ctx context.Context, country int) (dto.PhoneCode, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return dto.PhoneCode{}, err
	}
	defer conn.Release()
	var code string
	var codeID int
	row := conn.QueryRow(ctx, "SELECT id,phonecode FROM countries WHERE id = $1", country)
	err = row.Scan(&codeID, &code)
	if err != nil {
		return dto.PhoneCode{}, err
	}
	return dto.PhoneCode{
		ID:   codeID,
		Code: code,
	}, nil
}
