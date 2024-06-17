package identity

type Authority struct {
	authority string
}

func NewAuthority(authority string) Authority {
	return Authority{
		authority: authority,
	}
}

func (a Authority) Authority() string {
	return a.authority
}
