package find_recipient

type NameAndEmail struct {
	query string
	args  []interface{}
}

func ByNameAndEmail(name, email string) *NameAndEmail {
	return &NameAndEmail{
		query: string(`SELECT ID, IS_ENTERPRISE,SUBJECT_NAME,SUBJECT_THUMBNAIL,SUBJECT_EMAIL FROM RECIPIENTS
		 WHERE SUBJECT_NAME=$1 AND SUBJECT_EMAIL=$2`),
		args: []interface{}{name, email},
	}
}

func (e *NameAndEmail) Query() string {
	return e.query
}

func (e *NameAndEmail) Args() []interface{} {
	return e.args
}
