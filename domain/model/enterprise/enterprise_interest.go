package enterprise

import (
	"go-complaint/domain/model/common"

	"github.com/google/uuid"
)

type EnterpriseInterest struct {
	Id             uuid.UUID              `json:"id"`
	Name           string                 `json:"name"`
	LogoIMG        string                 `json:"logo_img"`
	BannerIMG      string                 `json:"banner_img"`
	Website        string                 `json:"website"`
	Email          string                 `json:"email"`
	Phone          string                 `json:"phone"`
	Address        common.AddressInterest `json:"address"`
	Industry       string                 `json:"industry"`
	FoundationDate string                 `json:"foundation_date"`
	OwnerId        uuid.UUID              `json:"owner_id"`
}
