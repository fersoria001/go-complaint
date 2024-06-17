package enterprise

import "go-complaint/domain"

type Industry struct {
	id       int
	name     string
	director domain.Director
}

func NewIndustry(id int, name string) (Industry, error) {
	var industry = &Industry{}
	industry.setID(id)
	err := industry.setName(name)
	if err != nil {
		return *industry, err
	}
	return *industry, nil
}

func NewIndustryWithDirector(director domain.Director) *Industry {
	return &Industry{
		director: director,
	}
}
func (i *Industry) Changed() {
	i.director.Changed(i)
}

func (i *Industry) setID(id int) {
	i.id = id
}

func (i *Industry) setName(name string) error {
	i.name = name
	return nil
}

func (i Industry) ID() int {
	return i.id
}

func (i Industry) Name() string {
	return i.name
}
