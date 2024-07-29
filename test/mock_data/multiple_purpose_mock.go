package mock_data

import (
	"go-complaint/domain/model/common"
	"time"
)

var (
	EmailVerificationToken = "emailVerificationToken"
	fixedDate              = time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC)
	CommonDate             = common.NewDate(fixedDate)
	Country                = common.NewCountry(11, "Argentina", "54")
	CountryState           = common.NewCountryState(3634, "San Juan")
	City                   = common.NewCity(644, "Albardón", "AR", -31.43722, -68.52556)
)
