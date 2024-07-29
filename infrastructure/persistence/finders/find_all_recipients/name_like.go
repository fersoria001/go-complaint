package find_all_recipients

type NameLike struct {
	query string
	args  []interface{}
}

func ByNameLike(query string) *NameLike {
	query = "%" + query + "%"
	return &NameLike{
		query: string(`SELECT ID, IS_ENTERPRISE, SUBJECT_NAME, SUBJECT_THUMBNAIL, SUBJECT_EMAIL
		 FROM RECIPIENTS WHERE LOWER(SUBJECT_NAME) LIKE LOWER($1) ORDER BY SUBJECT_NAME`),
		args: []interface{}{query},
	}
}

func (e *NameLike) Query() string {
	return e.query
}

func (e *NameLike) Args() []interface{} {
	return e.args
}
