package queries

import (
	"context"
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/dto"
	notificationsfindall "go-complaint/infrastructure/persistence/finders/notifications_findall"
	"go-complaint/infrastructure/persistence/repositories"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5"
)

type NotificationQuery struct {
	OwnerID string `json:"owner_id"`
}

func (notificationQuery NotificationQuery) Notifications(
	ctx context.Context,
) ([]dto.Notification, error) {
	if notificationQuery.OwnerID == "" {
		return nil, ErrBadRequest
	}
	segments := strings.Split(notificationQuery.OwnerID, "?")
	ids := make([]string, 0)
	for _, segment := range segments {
		if segment != "" {
			_, after, found := strings.Cut(segment, ":")
			if !found {
				continue
			}
			id := strings.Trim(after, " ")
			ids = append(ids, id)
		}
	}
	if len(ids) == 0 {
		return nil, ErrBadRequest
	}
	slice := make([]dto.Notification, 0)
	repository := repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository)
	for i := range ids {

		notifications, err := repository.FindAll(ctx, notificationsfindall.NewByOwnerID(ids[i]))
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {

				continue
			}

			return nil, err
		}

		notificationsSlice := notifications.ToSlice()
		for i := range notificationsSlice {
			slice = append(slice, dto.NewNotification(*notificationsSlice[i]))
		}
	}
	slices.SortStableFunc(slice, func(i, j dto.Notification) int {
		t1, _ := common.ParseDate(i.OccurredOn)
		t2, _ := common.ParseDate(j.OccurredOn)
		if t1.Before(t2) {
			return 1
		}
		if t1.After(t2) {
			return -1
		}
		return 0
	})
	return slice, nil
}
