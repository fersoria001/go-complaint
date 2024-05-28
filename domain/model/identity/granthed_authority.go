package identity

type GrantedAuthority interface {
	Authority() string
}
