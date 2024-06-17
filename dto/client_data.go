package dto

type ClientData struct {
	IP              string       `json:"ip"`
	Device          string       `json:"device"`
	Geolocalization Localization `json:"geolocalization"`
	LoginDate       string       `json:"login_date"`
}
