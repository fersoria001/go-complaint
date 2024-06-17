package common

type AddressDirector struct {
}

func NewAddressDirector(id string, a *Address) AddressDirector {
	return AddressDirector{}
}

func (ad AddressDirector) Show() {
}

func (ad AddressDirector) Create() {
}

func (ad AddressDirector) Changed(obj interface{}) {
}
