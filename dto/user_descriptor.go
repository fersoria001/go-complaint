package dto

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

type DeviceContextKey struct {
	name string
}

var DeviceCtxKey = DeviceContextKey{"device"}

type GeolocationContextKey struct {
	name string
}

var GeolocationCtxKey = GeolocationContextKey{"country"}

type IPContextKey struct {
	name string
}

var IPCtxKey = IPContextKey{"ip"}

// this looks more like a DTO
type UserDescriptor struct {
	Email       string    `json:"email"`
	FullName    string    `json:"full_name"`
	ProfileIMG  string    `json:"profile_img"`
	LoginDate   time.Time `json:"login_date"`
	IP          string    `json:"ip"`
	Device      string    `json:"device"`
	Geolocation string    `json:"geolocation"`
	RememberMe  bool      `json:"remember_me"`
	Role        string    `json:"role"`
	jwt.StandardClaims
}

func NewUserDescriptor(ctx context.Context, email, fullname, profileIMG string, rememberMe bool) (*UserDescriptor, error) {
	var (
		ip          string
		device      string
		geolocation string
	)

	// Check if IP value exists in context
	if value, ok := ctx.Value(IPCtxKey).(string); ok {
		ip = value
	} else {
		ip = ""
	}

	// Check if Device value exists in context
	if value, ok := ctx.Value(DeviceCtxKey).(string); ok {
		device = value
	} else {
		device = ""
	}

	// Check if Geolocation value exists in context
	if value, ok := ctx.Value(GeolocationCtxKey).(string); ok {
		geolocation = value
	} else {
		geolocation = ""
	}
	thisDate := time.Now()
	userDescriptor := &UserDescriptor{
		Email:       email,
		FullName:    fullname,
		ProfileIMG:  profileIMG,
		LoginDate:   thisDate,
		IP:          ip,
		Device:      device,
		Geolocation: geolocation,
		RememberMe:  rememberMe,
	}
	userDescriptor.IssuedAt = thisDate.Unix()
	userDescriptor.ExpiresAt = thisDate.Add(time.Hour * 24).Unix()
	return userDescriptor, nil
}
