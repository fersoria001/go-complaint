package application_services

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/repositories"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type contextKey string

var authorizationServiceInstance *AuthorizationApplicationService
var authorizationServiceOnce sync.Once

func AuthorizationApplicationServiceInstance() *AuthorizationApplicationService {
	authorizationServiceOnce.Do(func() {
		authorizationServiceInstance = NewAuthorizationApplicationService()
	})
	return authorizationServiceInstance
}

type AuthorizationApplicationService struct {
	deviceCtxKey      contextKey
	geolocationCtxKey contextKey
	ipCtxKey          contextKey
	tokenCtxKey       contextKey
	authCtxKey        contextKey
}

func NewAuthorizationApplicationService() *AuthorizationApplicationService {
	return &AuthorizationApplicationService{
		deviceCtxKey:      "device",
		geolocationCtxKey: "geolocation",
		ipCtxKey:          "ip",
		tokenCtxKey:       "jwt_token",
		authCtxKey:        "user_email",
	}
}
func (aas AuthorizationApplicationService) Credentials(ctx context.Context) (dto.UserDescriptor, error) {
	credentials, ok := ctx.Value(aas.authCtxKey).(dto.UserDescriptor)
	if !ok {
		return dto.UserDescriptor{}, &erros.UnauthorizedError{}
	}
	user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Get(ctx, credentials.Email)
	if err != nil {
		return credentials, err
	}
	credentials = dto.NewUserDescriptor(
		aas.ClientData(ctx),
		*user,
		false,
	)
	return credentials, nil
}

func (aas AuthorizationApplicationService) JWTToken(ctx context.Context) (string, error) {
	token, ok := ctx.Value(aas.tokenCtxKey).(string)
	if !ok {
		return "", &erros.UnauthorizedError{}
	}
	return token, nil
}

func (aas AuthorizationApplicationService) Authorize(
	ctx context.Context,
	jwtToken string,
) (context.Context, error) {
	claims, err := JWTApplicationServiceInstance().ParseUserDescriptor(jwtToken)
	if err != nil {
		claims, err := JWTApplicationServiceInstance().ParseConfirmationCode(jwtToken)
		if err != nil {
			return nil, err
		}
		authorizedCtx := context.WithValue(ctx, aas.authCtxKey, claims)
		authorizedCtx = context.WithValue(authorizedCtx, aas.tokenCtxKey, jwtToken)
		return authorizedCtx, nil
	}
	authorizedCtx := context.WithValue(ctx, aas.authCtxKey, claims)
	authorizedCtx = context.WithValue(authorizedCtx, aas.tokenCtxKey, jwtToken)
	return authorizedCtx, nil
}

func (aas AuthorizationApplicationService) ClientData(
	ctx context.Context,
) dto.ClientData {
	var (
		ip          string = ""
		device      string = ""
		geolocation        = [2]float64{0, 0}
	)
	if value, ok := ctx.Value(aas.ipCtxKey).(string); ok {
		ip = value
	}
	if value, ok := ctx.Value(aas.deviceCtxKey).(string); ok {
		device = value
	}
	if value, ok := ctx.Value(aas.geolocationCtxKey).([2]float64); ok {
		geolocation = value
	}
	thisDate := time.Now()
	commonDate := common.NewDate(thisDate).StringRepresentation()
	return dto.ClientData{
		IP:              ip,
		Device:          device,
		Geolocalization: dto.Localization{Latitude: geolocation[0], Longitude: geolocation[1]},
		LoginDate:       commonDate,
	}
}

type AccessLevel int

const (
	READ AccessLevel = iota + 1
	WRITE
	DELETE
)

func (aas AuthorizationApplicationService) ResourceAccess(
	ctx context.Context,
	resourceType string,
	resourceID string,
	accessLevel AccessLevel,
	requiredAuthorities ...string,
) (dto.UserDescriptor, error) {
	credentials, err := aas.Credentials(ctx)
	if err != nil {
		return credentials, ErrUnauthorized
	}
	if len(requiredAuthorities) == 0 {
		return credentials, nil
	}
	requiredAuthoritiesSet := mapset.NewSet[string]()
	for _, v := range requiredAuthorities {
		requiredAuthoritiesSet.Add(v)
	}
	switch resourceType {
	case "Complaint":
		parsedID, err := uuid.Parse(resourceID)
		if err != nil {
			return credentials, ErrBadRequest
		}
		dbComplaint, err := repositories.MapperRegistryInstance().Get("Complaint").(repositories.ComplaintRepository).Get(ctx, parsedID)
		if err != nil {
			return credentials, err
		}
		switch accessLevel {
		case WRITE:
			if dbComplaint.ReceiverID() != credentials.Email && dbComplaint.AuthorID() != credentials.Email {
				for _, v := range credentials.GrantedAuthorities {
					if (v.EnterpriseID == dbComplaint.ReceiverID() && requiredAuthoritiesSet.Contains(v.Authority)) ||
						(v.EnterpriseID == dbComplaint.AuthorID() && requiredAuthoritiesSet.Contains(v.Authority)) {
						return credentials, nil
					}
				}
			} else {
				return credentials, nil
			}
		}
	case "Feedback":
		parsedID, err := uuid.Parse(resourceID)
		if err != nil {
			return credentials, ErrBadRequest
		}
		dbFeedback, err := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository).Get(ctx, parsedID)
		if err != nil {
			return credentials, err
		}
		switch accessLevel {
		case WRITE:
			for _, v := range credentials.GrantedAuthorities {
				if v.EnterpriseID == dbFeedback.EnterpriseID() && requiredAuthoritiesSet.Contains(v.Authority) {
					return credentials, nil
				}
			}
			return credentials, ErrUnauthorized
		}
	case "Enterprise":
		enterprise, err := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository).Get(ctx, resourceID)
		if err != nil {
			return credentials, err
		}
		if enterprise.Owner() != credentials.Email {
			for _, v := range credentials.GrantedAuthorities {
				if v.EnterpriseID == resourceID && requiredAuthoritiesSet.Contains(v.Authority) {
					for emp := range enterprise.Employees().Iter() {
						if emp.Email() == credentials.Email && emp.Position().String() == v.Authority {
							return credentials, nil
						}
					}
				}
			}
		} else {
			return credentials, nil
		}
	default:
		if resourceID == "" {
			return credentials, nil
		}

		if resourceID != credentials.Email {
			for _, v := range credentials.GrantedAuthorities {

				if v.EnterpriseID == resourceID && requiredAuthoritiesSet.Contains(v.Authority) {
					return credentials, nil
				}
			}
		} else {
			return credentials, nil
		}
	}
	return credentials, ErrUnauthorized
}
