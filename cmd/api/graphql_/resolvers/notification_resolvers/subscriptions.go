package notification_resolvers

import (
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func NotificationsResolver(p graphql.ResolveParams) (interface{}, error) {
	nq := queries.NotificationQuery{
		OwnerID: p.Args["id"].(string),
	}
	notifications, err := nq.Notifications(p.Context)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
