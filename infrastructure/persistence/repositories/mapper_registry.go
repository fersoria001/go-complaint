package repositories

import (
	"go-complaint/infrastructure/persistence/datasource"
	"sync"
)

var mapperRegistry *MapperRegistry
var mapperRegistryOnce sync.Once

var staticMap = map[string]func() func(schema datasource.Schema) interface{}{
	"Country": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewCountryRepository(schema) }
	},
	"CountryState": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewCountryStateRepository(schema) }
	},
	"City": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewStateCitiesRepository(schema) }
	},
	"Address": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewAddressRepository(schema) }
	},
	"Person": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewPersonRepository(schema) }
	},
	"User": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewUserRepository(schema) }
	},
	"UserRole": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewUserRoleRepository(schema) }
	},
	"Enterprise": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewEnterpriseRepository(schema) }
	},
	"Employee": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewEmployeeRepository(schema) }
	},
	"Complaint": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewComplaintRepository(schema) }
	},
	"Reply": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewComplaintRepliesRepository(schema) }
	},
	"Feedback": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewFeedbackRepository(schema) }
	},
	"Answer": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewFeedbackAnswerRepository(schema) }
	},
	"FeedbackReply": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewFeedbackRepliesRepository(schema) }
	},
	"Review": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewFeedbackReviewRepository(schema) }
	},
	"ReplyReview": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewFeedbackReplyReviewRepository(schema) }
	},
	"Industry": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewIndustryRepository(schema) }
	},
	"Notification": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewNotificationRepository(schema) }
	},
	"Event": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewEventRepository(schema) }
	},
	"Chat": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewChatRepository(schema) }
	},
	"enterprise.Reply": func() func(schema datasource.Schema) interface{} {
		return func(schema datasource.Schema) interface{} { return NewChatRepliesRepository(schema) }
	},
}

func MapperRegistryInstance() *MapperRegistry {
	mapperRegistryOnce.Do(func() {
		mapperRegistry = NewMapperRegistry()
	})
	return mapperRegistry
}

type MapperRegistry struct {
	schema  datasource.Schema
	mappers map[string]func() func(schema datasource.Schema) interface{}
}

func NewMapperRegistry() *MapperRegistry {
	return &MapperRegistry{
		schema:  datasource.PublicSchema(),
		mappers: staticMap,
	}
}

func (mr *MapperRegistry) Register(
	key string,
	mapper func() func(schema datasource.Schema) interface{}) {
	mr.mappers[key] = mapper
}

func (mr *MapperRegistry) Get(key string) interface{} {
	if mr.mappers[key] == nil {
		return nil
	}
	builderMethod := mr.mappers[key]()
	return builderMethod(mr.schema)
}
