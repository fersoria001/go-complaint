package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
)

type PersonRepository struct {
	schema datasource.Schema
}

func NewPersonRepository(schema datasource.Schema) PersonRepository {
	return PersonRepository{schema: schema}
}

func (pr PersonRepository) Get(
	ctx context.Context,
	id uuid.UUID,
) (*identity.Person, error) {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	//
	selectQuery := string(`
		SELECT
		email,
		profile_img,
		genre,
		pronoun,
		first_name,
		last_name,
		birth_date,
		phone,
		address_id
		FROM person
		WHERE id = $1
		`)
	var (
		email      string
		profileIMG string
		gender     string
		pronoun    string
		firstName  string
		lastName   string
		birthDate  string
		phone      string
		addressID  uuid.UUID
	)
	err = conn.QueryRow(
		ctx,
		selectQuery,
		id,
	).Scan(
		&email,
		&profileIMG,
		&gender,
		&pronoun,
		&firstName,
		&lastName,
		&birthDate,
		&phone,
		&addressID,
	)
	if err != nil {
		return nil, err
	}
	mapper := MapperRegistryInstance().Get("Address")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	addressRepository, ok := mapper.(AddressRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	address, err := addressRepository.Get(ctx, addressID)
	if err != nil {
		return nil, err
	}
	stringBirthDate, err := common.NewDateFromString(birthDate)
	if err != nil {
		return nil, err
	}
	person, err := identity.NewPerson(
		id,
		email,
		profileIMG,
		gender,
		pronoun,
		firstName,
		lastName,
		phone,
		stringBirthDate,
		address,
	)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return person, nil
}

func (pr PersonRepository) Save(
	ctx context.Context,
	person identity.Person,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("Address")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	addressRepository, ok := mapper.(AddressRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	insertCommand := string(`
		INSERT INTO 
		person (
			id,
			email,
			profile_img,
			genre,
			pronoun,
			first_name,
			last_name,
			birth_date,
			phone,
			address_id
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	var (
		id         uuid.UUID = person.Id()
		email      string    = person.Email()
		profileIMG string    = person.ProfileIMG()
		gender     string    = person.Genre()
		pronoun    string    = person.Pronoun()
		firstName  string    = person.FirstName()
		lastName   string    = person.LastName()
		birthDate  string    = person.BirthDate().StringRepresentation()
		phone      string    = person.Phone()
		addressID  uuid.UUID = person.Address().ID()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&email,
		&profileIMG,
		&gender,
		&pronoun,
		&firstName,
		&lastName,
		&birthDate,
		&phone,
		&addressID,
	)
	if err != nil {
		return err
	}
	err = addressRepository.Save(ctx, person.Address())
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (pr PersonRepository) Remove(
	ctx context.Context,
	id uuid.UUID,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "DELETE FROM PERSON WHERE PERSON.ID=$1", &id)
	if err != nil {
		return err
	}
	addressRepository := MapperRegistryInstance().Get("Address").(AddressRepository)
	err = addressRepository.Remove(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (pr PersonRepository) Update(
	ctx context.Context,
	person identity.Person,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("Address")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	addressRepository, ok := mapper.(AddressRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	insertCommand := string(`
		UPDATE person
		SET profile_img=$2,
			genre=$3 ,
			pronoun=$4 ,
			first_name=$5,
			last_name=$6,
			birth_date=$7,
			phone=$8,
			address_id=$9
		WHERE id = $1`)
	var (
		id         uuid.UUID = person.Id()
		profileIMG string    = person.ProfileIMG()
		gender     string    = person.Genre()
		pronoun    string    = person.Pronoun()
		firstName  string    = person.FirstName()
		lastName   string    = person.LastName()
		birthDate  string    = person.BirthDate().StringRepresentation()
		phone      string    = person.Phone()
		addressID  uuid.UUID = person.Address().ID()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&profileIMG,
		&gender,
		&pronoun,
		&firstName,
		&lastName,
		&birthDate,
		&phone,
		&addressID,
	)
	if err != nil {
		return err
	}
	err = addressRepository.Update(ctx, person.Address())
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}
