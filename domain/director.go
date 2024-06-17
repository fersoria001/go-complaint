package domain

type Director interface {
	Show() interface{}
	Create()
	Changed(obj interface{}) error
}
