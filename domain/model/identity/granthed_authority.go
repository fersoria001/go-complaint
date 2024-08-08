package identity

type GrantedAuthority interface {
	Principal() string
	Authority() string
}
