package models

import (
	"time"
)

type Organization struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrganization(name string, id int) *Organization {
	return &Organization{
		Id:   id,
		Name: name,
	}
}
