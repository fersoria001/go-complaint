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
	personEmail string,
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
		gender,
		pronoun,
		first_name,
		last_name,
		birth_date,
		phone,
		address_id
		FROM person
		WHERE email = $1
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
		personEmail,
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
			email,
			profile_img,
			gender,
			pronoun,
			first_name,
			last_name,
			birth_date,
			phone,
			address_id
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	var (
		email      string    = person.Email()
		profileIMG string    = person.ProfileIMG()
		gender     string    = person.Gender()
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
			gender=$3 ,
			pronoun=$4 ,
			first_name=$5,
			last_name=$6,
			birth_date=$7,
			phone=$8,
			address_id=$9
		WHERE email = $1`)
	var (
		email      string    = person.Email()
		profileIMG string    = person.ProfileIMG()
		gender     string    = person.Gender()
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
	err = addressRepository.Update(ctx, person.Address())
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}
