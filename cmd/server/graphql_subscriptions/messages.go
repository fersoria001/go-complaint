package graphql_subscriptions

type SubscriptionsMessageType int

const (
	UNSUPPORTED SubscriptionsMessageType = iota
	CONNECTIONACK
	DATA
)

func (m SubscriptionsMessageType) String() string {
	switch m {
	case DATA:
		return "data"
	case CONNECTIONACK:
		return "connection_ack"
	default:
		return "not_supported"
	}
}

type ConnectionACKMessage struct {
	Type        string `json:"type"`
	OperationID string `json:"operation_id"`
	Payload     struct {
		Query          string `json:"query"`
		SubscriptionID string `json:"subscription_id"`
		Token          string `json:"token"`
		EnterpriseID   string `json:"enterprise_id"`
	} `json:"payload"`
}

type DataMessage struct {
	Type        string      `json:"type"`
	OperationID string      `json:"operation_id"`
	Payload     interface{} `json:"payload"`
}
type Message struct {
	SubscriptionID int    `json:"subscription_id"`
	Token          string `json:"token"`
	EnterpriseID   string `json:"enterprise_id"`
	Type           string `json:"type"`
	Payload        []byte `json:"content"`
}
