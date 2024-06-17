package repositories

type StatementSource interface {
	Query() string
	Args() []interface{}
}
