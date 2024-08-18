package commands

import (
	"context"
	"errors"
	"fmt"
	"go-complaint/application"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

type LogEnterpriseActivityCommand struct {
	UserId         string `json:"userId"`
	ActivityId     string `json:"activityId"`
	EnterpriseId   string `json:"enterpriseId"`
	EnterpriseName string `json:"enterpriseName"`
	ActivityType   string `json:"activityType"`
}

func NewLogEnterpriseActivityCommand(userId, activityId, enterpriseId, enterpriseName, activityType string) *LogEnterpriseActivityCommand {
	return &LogEnterpriseActivityCommand{
		UserId:         userId,
		ActivityId:     activityId,
		EnterpriseId:   enterpriseId,
		EnterpriseName: enterpriseName,
		ActivityType:   activityType,
	}
}

func (c LogEnterpriseActivityCommand) Execute(ctx context.Context) error {
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("EnterpriseActivity").(repositories.EnterpriseActivityRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	recipientRepository := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	userId, err := uuid.Parse(c.UserId)
	if err != nil {
		return err
	}
	activityId, err := uuid.Parse(c.ActivityId)
	if err != nil {
		return err
	}
	enterpriseId, err := uuid.Parse(c.EnterpriseId)
	if err != nil {
		return err
	}
	activityType := enterprise.ParseActivityType(c.ActivityType)
	if activityType < 0 {
		return fmt.Errorf("unexpected activity type at log enterprise activity command")
	}

	user, err := recipientRepository.Get(ctx, userId)
	if err != nil {
		return err
	}
	ea := enterprise.NewEnterpriseActivity(
		uuid.New(), activityId, enterpriseId, *user, c.EnterpriseName, time.Now(), activityType,
	)
	err = r.Save(ctx, *ea)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" &&
				pgErr.ConstraintName == "enterprise_activity_user_id_activity_id_key" {
				return ErrEnterpriseActivityAlreadyExists
			}
		}
		return err
	}
	svc := application.ApplicationMessagePublisherInstance()
	svc.Publish(application.NewApplicationMessage(
		user.Id().String(),
		"enterpriseActivity",
		*dto.NewEnterpriseActivity(*ea),
	))
	svc.Publish(application.NewApplicationMessage(
		c.EnterpriseName,
		"enterpriseActivity",
		*dto.NewEnterpriseActivity(*ea),
	))
	return nil
}
