package common

type Country struct {
	id        int
	name      string
	phonecode string
}

func NewCountry(id int, name, phonecode string) Country {
	return Country{id: id, name: name, phonecode: phonecode}
}

func (c Country) ID() int {
	return c.id
}

func (c Country) Name() string {
	return c.name
}

func (c Country) PhoneCode() string {
	return c.phonecode
}
