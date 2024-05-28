package application

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"log"
	"reflect"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

type EnterpriseService struct {
	eventStore           *NotificationService
	enterpriseRepository *repositories.EnterpriseRepository
	employeeRepository   *repositories.EmployeeRepository
	identityService      *IdentityService
}

func NewEnterpriseService(
	enterpriseRepository *repositories.EnterpriseRepository,
	identityService *IdentityService,
	employeeRepository *repositories.EmployeeRepository,
	eventStore *NotificationService,
) *EnterpriseService {
	return &EnterpriseService{
		enterpriseRepository: enterpriseRepository,
		identityService:      identityService,
		employeeRepository:   employeeRepository,
		eventStore:           eventStore,
	}
}

// remember to compare enterpriseName := strings.Split(managerID, "-")[0]
// of both managerID and employeeID guess this belong to owner
func (enterpriseService *EnterpriseService) FireEmployee(
	ctx context.Context,
	enterpriseName string,
	employeeID string) error {
	var (
		employee     *enterprise.Employee
		enterpriseDB *enterprise.Enterprise
		err          error
	)

	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.EmployeeFired); ok {
				fmt.Println("Employee Fired Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeFired{})
		},
	})

	splitID := strings.Split(employeeID, "-")
	if len(splitID) != 4 {
		return &erros.ValidationError{Expected: "invalid employeeID"}
	}
	if splitID[0] != enterpriseName {
		return &erros.ValidationError{Expected: "an employee who belongs to this enterprise"}
	}
	employee, err = enterpriseService.employeeRepository.Get(ctx, employeeID)
	if err != nil {
		return err
	}
	enterpriseDB, err = enterpriseService.enterpriseRepository.Get(ctx, enterpriseName)
	if err != nil {
		return err
	}
	err = enterpriseDB.Owner().FireEmployee(ctx, employee)
	if err != nil {
		return err
	}
	err = enterpriseService.identityService.RemoveRole(
		ctx,
		employee.Email(),
		employee.Position().String(),
		enterpriseName,
	)
	if err != nil {
		return err
	}
	err = enterpriseService.employeeRepository.Remove(ctx, employeeID)
	if err != nil {
		return err
	}
	return nil
}

// this should be owner only for now because there is only two roles
func (enterpriseService *EnterpriseService) PromoteEmployee(
	ctx context.Context,
	employeeID,
	position string) error {
	var (
		enterpriseDB *enterprise.Enterprise
		employee     *enterprise.Employee
		err          error
	)
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.EmployeePromoted); ok {
				fmt.Println("Employee Promoted Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeePromoted{})
		},
	})
	splitID := strings.Split(employeeID, "-")
	if len(splitID) != 4 {
		return &erros.ValidationError{Expected: "invalid employeeID"}
	}
	enterpriseName := splitID[0]
	enterpriseDB, err = enterpriseService.enterpriseRepository.Get(ctx, enterpriseName)
	if err != nil {
		return err
	}
	employee, err = enterpriseService.employeeRepository.Get(ctx, employeeID)
	if err != nil {
		return err
	}
	lastPosition := employee.Position().String()
	err = enterpriseDB.Owner().PromoteEmployee(ctx, enterpriseName, employee, position)
	if err != nil {
		return err
	}
	err = enterpriseService.identityService.ChangeRole(
		ctx,
		employee.Email(),
		enterpriseName,
		lastPosition,
		position,
	)
	if err != nil {
		err = enterpriseService.identityService.AddNewRole(
			ctx,
			employee.Email(),
			position,
			enterpriseName,
		)
		if err != nil {
			return err
		}
	}
	return enterpriseService.employeeRepository.Update(ctx, employee)
}

func (enterpriseService *EnterpriseService) CancelHiring(
	ctx context.Context,
	enterpriseName string,
	employeeID string) error {
	var (
		employee     *enterprise.Employee
		enterpriseDB *enterprise.Enterprise
		err          error
	)

	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.HiringProccessCanceled); ok {
				fmt.Println("Employee Fired Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringProccessCanceled{})
		},
	})

	splitID := strings.Split(employeeID, "-")
	if len(splitID) != 4 {
		return &erros.ValidationError{Expected: "invalid employeeID"}
	}
	if splitID[0] != enterpriseName {
		return &erros.ValidationError{Expected: "an employee who belongs to this enterprise"}
	}
	employee, err = enterpriseService.employeeRepository.Get(ctx, employeeID)
	if err != nil {
		return err
	}
	enterpriseDB, err = enterpriseService.enterpriseRepository.Get(ctx, enterpriseName)
	if err != nil {
		return err
	}
	err = enterpriseDB.Owner().CancelHiring(ctx, employee)
	if err != nil {
		return err
	}
	err = enterpriseService.identityService.RemoveRole(
		ctx,
		employee.Email(),
		employee.Position().String(),
		enterpriseName,
	)
	if err != nil {
		return err
	}
	err = enterpriseService.employeeRepository.Remove(ctx, employeeID)
	if err != nil {
		return err
	}
	return nil
}

/*
	this will approve an employee that has been prehired

with the flag ApprovedHiring set to false
*/
func (enterpriseService *EnterpriseService) ApproveEmployee(
	ctx context.Context,
	enterpriseName,
	employeeID string) error {
	var (
		enterpriseDB *enterprise.Enterprise
		employee     *enterprise.Employee
		owner        *enterprise.Owner
		err          error
	)
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.EmployeeHired); ok {
				fmt.Println("Employee Hired Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeHired{})
		},
	})

	employee, err = enterpriseService.employeeRepository.Get(ctx, employeeID)
	if err != nil {
		return err
	}
	enterpriseDB, err = enterpriseService.enterpriseRepository.Get(ctx, enterpriseName)
	if err != nil {
		return err
	}
	owner = enterpriseDB.Owner()
	err = owner.ApproveHiring(ctx, employee, enterpriseName)
	if err != nil {
		return err
	}
	err = enterpriseService.employeeRepository.Update(ctx, employee)
	if err != nil {
		return err
	}
	err = enterpriseService.identityService.AddNewRole(
		ctx,
		employee.Email(),
		employee.Position().String(),
		enterpriseName)
	if err != nil {
		return err
	}
	return nil
}
func (enterpriseService *EnterpriseService) AcceptHiringInvitation(ctx context.Context,
	invitationEventID string,
) error {
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			handlerCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if _, ok := event.(*enterprise.EmployeeWaitingForApproval); ok {
				fmt.Println("Accept Hiring Invitation Handler, handle the removal of invitationSent")
				eventStore := repositories.NewEventRepository(datasource.EventSchema())
				eventService := NewNotificationService(eventStore)
				return eventService.MoveToLog(handlerCtx, invitationEventID)
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeWaitingForApproval{})
		},
	})
	storedEvent, err := enterpriseService.eventStore.ProvideFromStore(ctx, invitationEventID)
	if err != nil {
		return err
	}
	var invitation *enterprise.HiringInvitationSent
	err = json.Unmarshal(storedEvent.EventBody, &invitation)
	if err != nil {
		return err
	}
	ep, err := enterpriseService.Enterprise(ctx, invitation.EnterpriseID())
	if err != nil {
		return err
	}
	owner := enterprise.NewOwner(ep.OwnerID)
	employeeID := enterpriseService.NewEmployeeID(
		invitation.EnterpriseID(),
		ep.OwnerID,
		invitation.Email(),
	)
	employeeWaitingForApproval, err := owner.AcceptHiringInvitation(
		ctx,
		storedEvent.EventId,
		employeeID,
		invitation.ProfileIMG(),
		invitation.FirstName(),
		invitation.LastName(),
		invitation.Email(),
		invitation.Phone(),
		invitation.Age(),
		invitation.ProposalPosition(),
	)
	if err != nil {
		return err
	}
	err = enterpriseService.employeeRepository.Save(ctx, employeeWaitingForApproval)
	if err != nil {
		return err
	}
	return nil
}

// func (enterpriseService *EnterpriseService) HireInvitedUser(
// 	ctx context.Context,
// 	ownerID,
// 	eventID string) error {
// 	retrievedEvent, err := enterpriseService.eventStore.Get(ctx, eventID)
// 	if err != nil {
// 		return err
// 	}
// 	var hiringInvitationSent *enterprise.HiringInvitationSent
// 	err = json.Unmarshal(retrievedEvent.EventBody, &hiringInvitationSent)
// 	if err != nil {
// 		return err
// 	}
// 	enterpriseDB, err := enterpriseService.enterpriseRepository.Get(
// 		ctx,
// 		hiringInvitationSent.EnterpriseID(),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	if enterpriseDB.OwnerID() != ownerID {
// 		return &erros.UnauthorizedError{}
// 	}
// 	employee, err := enterpriseDB.Owner().Hire(ctx, hiringInvitationSent)
// 	if err != nil {
// 		return err
// 	}
// 	err = enterpriseService.employeeRepository.Save(ctx, employee)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (enterprisesService *EnterpriseService) InviteToProject(
	ctx context.Context,
	enterpriseName string,
	ownerID,
	userEmail string,
	position string) error {
	newPosition, err := enterprise.ParsePosition(position)
	if err != nil {
		return err
	}
	enterpriseDB, err := enterprisesService.enterpriseRepository.Get(
		ctx,
		enterpriseName,
	)
	if err != nil {
		return err
	}
	//this should be done before executing this method
	//to avoid unnecessary calls to the database
	//it's only one call it can be done twice
	//at the cmd level and here
	if enterpriseDB.OwnerID() != ownerID {
		return &erros.UnauthorizedError{}
	}

	employeedIn, _ := enterprisesService.employeeRepository.FindByEmail(ctx, userEmail)
	if employeedIn != nil {
		for ep := range employeedIn.Iter() {
			nameSegmentOfID := strings.Split(ep.ID(), "-")[0]
			if nameSegmentOfID == enterpriseName {
				return &erros.AlreadyExistsError{}
			}
		}
	}
	employeedIn.Clear()
	user, err := enterprisesService.identityService.User(ctx, userEmail)
	if err != nil {
		return err
	}
	userAge := time.Now().Year() - user.Person().BirthDate().Date().Year()
	owner := enterpriseDB.Owner()
	newEmployeeID := enterprisesService.NewEmployeeID(
		enterpriseName,
		ownerID,
		userEmail,
	)
	err = owner.SendHiringInvitation(
		ctx,
		enterpriseName,
		newEmployeeID,
		newPosition,
		user.ProfileIMG(),
		user.Person().FirstName(),
		user.Person().LastName(),
		user.Email(),
		user.Person().Phone(),
		userAge)
	if err != nil {
		return err
	}
	return nil
}
func (enterpriseService *EnterpriseService) ProvideEmployeeOffices(
	ctx context.Context,
	userID string,
) ([]*dto.Office, error) {
	employs, err := enterpriseService.employeeRepository.FindByEmail(ctx, userID)
	if err != nil {
		return nil, err
	}
	offices := map[string]*dto.Office{}
	enterprisesIDS := mapset.NewSet[string]()
	for empl := range employs.Iter() {
		enterpriseID := strings.Split(empl.ID(), "-")
		if len(enterpriseID) != 4 {
			return nil, err
		}
		enterprisesIDS.Add(enterpriseID[0])
		offices[enterpriseID[0]] = &dto.Office{
			EmployeeID:        empl.ID(),
			EmployeeFirstName: empl.FirstName(),
			EmployeePosition:  empl.Position().String(),
			EnterpriseName:    enterpriseID[0],
		}
	}
	enterprises, err := enterpriseService.enterpriseRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range enterprises.ToSlice() {
		if !enterprisesIDS.Contains(item.Name()) {
			enterprises.Remove(item)
		}
	}
	result := make([]*dto.Office, 0)
	for item := range enterprises.Iter() {
		office := offices[item.Name()]
		office.EnterpriseLogoIMG = item.LogoIMG()
		office.EnterpriseWebsite = item.Website()
		office.EnterpriseIndustry = item.Industry().Name()
		office.EnterpriseAddress.Country = item.Address().Country()
		office.EnterpriseAddress.County = item.Address().County()
		office.EnterpriseAddress.City = item.Address().City()
		owner, err := enterpriseService.identityService.repository.Get(
			ctx,
			item.OwnerID(),
		)
		if err != nil {
			return nil, err
		}
		office.OwnerFullName = owner.Person().FullName()
		result = append(result, office)
	}
	return result, nil
}

// AvailableUsers(ctx context.Context) ([]*dto.UserDTO, error)
func (enterpriseService *EnterpriseService) AvailableUsers(
	ctx context.Context,
	ownerID,
	query,
	enterpriseID string,
	limit, offset int) ([]*dto.User, int, error) {
	actualEmployees, err := enterpriseService.employeeRepository.GetAll(ctx)
	if err != nil {
		return nil, 0, err
	}
	employeesID := make([]string, 0)
	//all users that its id is not the email from employee.id[0]
	for employee := range actualEmployees.Iter() {
		segments := strings.Split(employee.ID(), "-")
		if segments[0] != enterpriseID {
			employeesID = append(employeesID, employee.Email())
		}
	}
	employeesID = append(employeesID, ownerID)
	employables, count, err := enterpriseService.identityService.PossibleEmployees(ctx, query, employeesID, limit, offset)
	if err != nil {
		return nil, count, err
	}
	results := make([]*dto.User, 0)
	for _, employee := range employables {
		results = append(results, dto.NewUser(employee))
	}
	return results, count, nil
}
func (enterpriseService *EnterpriseService) Enterprise(
	ctx context.Context,
	enterpriseID string) (*dto.Enterprise, error) {
	ep, err := enterpriseService.enterpriseRepository.Get(ctx, enterpriseID)
	if err != nil {
		log.Printf("EnterpriseService:Enterprise with id %s: %s\n ", enterpriseID, err.Error())
		return nil, err
	}
	enterpriseDto := dto.NewEnterprise(ep)
	return enterpriseDto, nil
}

// think about pagination, return count
func (enterpriseService *EnterpriseService) ProvideOwnerEnterprises(
	ctx context.Context,
	ownerID string) (mapset.Set[dto.Enterprise], error) {
	var enterprisesDTOS mapset.Set[dto.Enterprise] = mapset.NewSet[dto.Enterprise]()
	enterprises, err := enterpriseService.enterpriseRepository.GetAll(ctx)

	if err != nil {
		return nil, err
	}
	for ep := range enterprises.Iter() {
		if ep.OwnerID() != ownerID {
			log.Printf("OwnerID: %s, Enterprise OwnerID: %s\n",
				ownerID, ep.OwnerID())
			enterprises.Remove(ep)
		}
	}
	// if enterprises.Cardinality() == 0 {
	// 	return nil, &erros.ValueNotFoundError{}
	// }
	for ep := range enterprises.Iter() {
		enterprisesDTOS.Add(*dto.NewEnterprise(ep))
	}
	return enterprisesDTOS, nil
}

func (enterpriseService *EnterpriseService) UpdateEnterprise(
	ctx context.Context,
	id string,
	logoIMG,
	website,
	email,
	phone,
	country,
	county,
	city string) error {
	var (
		ep  *enterprise.Enterprise
		err error
	)
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if _, ok := event.(*enterprise.EnterpriseUpdated); ok {
					fmt.Println("Enterprise Updated Handler")
					return nil
				}
				return &erros.ValueNotFoundError{}
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&enterprise.EnterpriseUpdated{})
			},
		})

	ep, err = enterpriseService.enterpriseRepository.Get(ctx, id)
	if err != nil {
		return err
	}
	if logoIMG != "" {
		err = ep.ChangeLogoIMG(ctx, logoIMG)
		if err != nil {
			return err
		}
	}
	if website != "" {
		err = ep.ChangeWebsite(ctx, website)
		if err != nil {
			return err
		}
	}
	if email != "" {
		err = ep.ChangeEmail(ctx, email)
		if err != nil {
			return err
		}
	}
	if phone != "" {
		err = ep.ChangePhone(ctx, phone)
		if err != nil {
			return err
		}
	}
	if country != "" || county != "" || city != "" {
		err = ep.ChangeAddress(ctx, country, county, city)
		if err != nil {
			return err
		}
	}
	err = enterpriseService.enterpriseRepository.Update(ctx, ep)
	if err != nil {
		return err
	}

	return nil
}

func (enterpriseService *EnterpriseService) OwnerOfEnterprise(ctx context.Context, id string) (mapset.Set[*enterprise.Enterprise], error) {
	return enterpriseService.enterpriseRepository.FindByOwnerID(ctx, id)
}

/*
	preconditions: ownerID must represent a valid user

enterpriseName, industry and website must be already validated
and unique
*/
func (enterpriseService *EnterpriseService) CreateEnterprise(
	ctx context.Context,
	ownerID string,
	enterpriseName,
	website,
	email,
	phone,
	country,
	county,
	city,
	industryName,
	foundationDate string) error {
	var (
		publisher = domain.DomainEventPublisherInstance()
		ep        *enterprise.Enterprise
		err       error
	)
	defaultLogoIMG := "/default.jpg"
	defaultBannerIMG := "/enterprise-banner.jpg"
	publisher.Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.EnterpriseCreated); ok {
				fmt.Println("Enterprise Created Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EnterpriseCreated{})
		},
	})
	owner, err := enterpriseService.identityService.User(ctx, ownerID)
	if err != nil {
		return err
	}
	ep, err = enterprise.CreateEnterprise(ctx,
		ownerID,
		enterpriseName,
		defaultLogoIMG,
		defaultBannerIMG,
		website,
		email,
		phone,
		country,
		county,
		city,
		industryName,
		foundationDate)
	if err != nil {
		return err
	}
	err = enterpriseService.enterpriseRepository.Save(ctx, ep)
	if err != nil {
		return err
	}
	err = enterpriseService.identityService.AddNewRole(ctx,
		owner.Email(), "OWNER", enterpriseName)
	if err != nil {
		return err
	}
	return nil
}

func (owns *EnterpriseService) IsAvailable(
	ctx context.Context,
	enterpriseName string) error {
	_, err := owns.enterpriseRepository.Get(ctx, enterpriseName)
	if err != nil {
		return nil
	}
	return &erros.AlreadyExistsError{}
}

func (enterpriseService *EnterpriseService) NewEmployeeID(
	enterpriseName,
	managerEmail,
	userID string) string {
	return enterpriseName + "-" + managerEmail + "-" + time.Now().Format("02/01/2006") + "-" + userID
}
