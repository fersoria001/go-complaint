package identity

type Authority struct {
	principal string
	authority string
}

func NewAuthority(principal, authority string) Authority {
	return Authority{
		principal: principal,
		authority: authority,
	}
}

func (a Authority) Principal() string {
	return a.principal
}

func (a Authority) Authority() string {
	return a.authority
}
