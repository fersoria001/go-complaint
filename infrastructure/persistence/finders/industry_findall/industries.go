package industryfindall

type Industries struct {
	query string
	args  []interface{}
}

func NewIndustries() *Industries {
	return &Industries{
		query: string(
			`SELECT 
		industry_id,
		name
		FROM industry`,
		),
		args: []interface{}{},
	}
}

func (e *Industries) Query() string {
	return e.query
}

func (e *Industries) Args() []interface{} {
	return e.args
}
