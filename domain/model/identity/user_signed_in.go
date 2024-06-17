package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type UserSignedIn struct {
	userID           string
	confirmationCode int
	ip               string
	device           string
	latitude         float64
	longitude        float64
	occurredOn       time.Time
}

func NewUserSignedIn(
	userID string,
	confirmationCode int,
	ip string,
	device string,
	latitude float64,
	longitude float64,
	occurredOn time.Time,
) *UserSignedIn {
	return &UserSignedIn{
		userID:           userID,
		confirmationCode: confirmationCode,
		ip:               ip,
		device:           device,
		latitude:         latitude,
		longitude:        longitude,
		occurredOn:       occurredOn,
	}
}

func (userSignedIn UserSignedIn) UserID() string {
	return userSignedIn.userID
}

func (userSignedIn UserSignedIn) ConfirmationCode() int {
	return userSignedIn.confirmationCode
}

func (userSignedIn UserSignedIn) IP() string {
	return userSignedIn.ip
}

func (userSignedIn UserSignedIn) Device() string {
	return userSignedIn.device
}

func (userSignedIn UserSignedIn) Latitude() float64 {
	return userSignedIn.latitude
}

func (userSignedIn UserSignedIn) Longitude() float64 {
	return userSignedIn.longitude
}

func (userSignedIn UserSignedIn) OccurredOn() time.Time {
	return userSignedIn.occurredOn
}

func (userSignedIn *UserSignedIn) MarshalJSON() ([]byte, error) {
	commonDate := common.NewDate(userSignedIn.occurredOn)
	return json.Marshal(map[string]interface{}{
		"user_id":           userSignedIn.userID,
		"confirmation_code": userSignedIn.confirmationCode,
		"ip":                userSignedIn.ip,
		"device":            userSignedIn.device,
		"latitude":          userSignedIn.latitude,
		"longitude":         userSignedIn.longitude,
		"occurred_on":       commonDate.StringRepresentation(),
	})
}

func (userSignedIn *UserSignedIn) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	userSignedIn.userID = v["user_id"].(string)
	userSignedIn.confirmationCode = v["confirmation_code"].(int)
	userSignedIn.ip = v["ip"].(string)
	userSignedIn.device = v["device"].(string)
	userSignedIn.latitude = v["latitude"].(float64)
	userSignedIn.longitude = v["longitude"].(float64)
	commonDate, err := common.ParseDate(v["occurred_on"].(string))
	if err != nil {
		return err
	}
	userSignedIn.occurredOn = commonDate
	return nil
}
