package common

type Localization struct {
	latitude  float64
	longitude float64
}

func NewLocalization(latitude float64, longitude float64) Localization {
	return Localization{
		latitude:  latitude,
		longitude: longitude,
	}
}

func (l Localization) Latitude() float64 {
	return l.latitude
}

func (l Localization) Longitude() float64 {
	return l.longitude
}
