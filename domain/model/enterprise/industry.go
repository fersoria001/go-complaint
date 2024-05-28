package enterprise

import (
	"go-complaint/erros"
)

type Industry struct {
	id   int
	name string
}

func NewIndustry(id int, name string) (Industry, error) {
	var industry = Industry{}
	industry.setID(id)
	err := industry.setName(name)
	if err != nil {
		return industry, err
	}
	return industry, nil
}

func (i *Industry) setID(id int) {
	i.id = id
}

func (i *Industry) setName(name string) error {
	if name == "" {
		return &erros.NullValueError{}
	}
	i.name = name
	return nil
}

func (i Industry) ID() int {
	return i.id
}

func (i Industry) Name() string {
	return i.name
}
