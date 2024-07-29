package enterprise_test

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/test/mock_data"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnterprise_NewEnterprise(t *testing.T) {
	for _, v := range mock_data.NewEnterprises {
		industry, err := enterprise.NewIndustry(v.Industry.Id, v.Industry.Name)
		assert.Nil(t, err)
		e, err := enterprise.NewEnterprise(
			v.Id,
			v.OwnerId,
			v.Name,
			v.LogoImg,
			v.BannerImg,
			v.Website,
			v.Email,
			v.Phone,
			v.Address,
			industry,
			v.RegisterAt,
			v.UpdatedAt,
			v.FoundationDate,
			v.Employees,
		)
		assert.Nil(t, err)
		assert.NotNil(t, e)
		assert.Equal(t, v.Id, e.Id())
		assert.Equal(t, v.OwnerId, e.OwnerId())
		assert.Equal(t, v.Name, e.Name())
		assert.Equal(t, v.LogoImg, e.LogoIMG())
		assert.Equal(t, v.BannerImg, e.BannerIMG())
		assert.Equal(t, v.Website, e.Website())
		assert.Equal(t, v.Email, e.Email())
		assert.Equal(t, v.Phone, e.Phone())
		assert.Equal(t, v.Address, e.Address())
		assert.Equal(t, industry, e.Industry())
		assert.Equal(t, v.UpdatedAt, e.UpdatedAt())
		assert.Equal(t, v.FoundationDate, e.FoundationDate())
		assert.Equal(t, v.Employees.Cardinality(), e.Employees().Cardinality())
	}
}

func TestEnterprise_CreateEnterprise(t *testing.T) {
	ctx := context.Background()
	for _, v := range mock_data.NewEnterprises {
		industry, err := enterprise.NewIndustry(v.Industry.Id, v.Industry.Name)
		assert.Nil(t, err)
		c := 0
		domain.DomainEventPublisherInstance().Subscribe(
			domain.DomainEventSubscriber{
				HandleEvent: func(event domain.DomainEvent) error {
					if _, ok := event.(*enterprise.EnterpriseCreated); !ok {
						t.Fatalf("Incorrect EventType in Subscriber")
					}
					c++
					return nil
				},
				SubscribedToEventType: func() reflect.Type {
					return reflect.TypeOf(&enterprise.EnterpriseCreated{})
				},
			},
		)
		e, err := enterprise.CreateEnterprise(
			ctx,
			v.Id,
			v.OwnerId,
			v.Name,
			v.LogoImg,
			v.BannerImg,
			v.Website,
			v.Email,
			v.Phone,
			v.FoundationDate,
			industry,
			v.Address,
		)
		assert.Nil(t, err)
		assert.NotNil(t, e)
		assert.Equal(t, v.Id, e.Id())
		assert.Equal(t, v.OwnerId, e.OwnerId())
		assert.Equal(t, v.Name, e.Name())
		assert.Equal(t, v.LogoImg, e.LogoIMG())
		assert.Equal(t, v.BannerImg, e.BannerIMG())
		assert.Equal(t, v.Website, e.Website())
		assert.Equal(t, v.Email, e.Email())
		assert.Equal(t, v.Phone, e.Phone())
		assert.Equal(t, v.Address, e.Address())
		assert.Equal(t, industry, e.Industry())
		assert.Equal(t, v.FoundationDate.StringRepresentation(), e.FoundationDate().StringRepresentation())
		assert.Equal(t, v.Employees.Cardinality(), e.Employees().Cardinality())
		assert.Equal(t, 1, c)
	}
}
