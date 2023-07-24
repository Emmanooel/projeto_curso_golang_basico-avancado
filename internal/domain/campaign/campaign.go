package campaign

import (
	"errors"
	"time"
)

type Contact struct {
	Email string
}

type Campaign struct {
	Id       string
	Name     string
	Content  string
	CreateOn time.Time
	Contacts []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("contacts is required")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		Id:       "ciqvh52navij03c99p5g",
		Name:     name,
		Content:  content,
		CreateOn: time.Now(),
		Contacts: contacts,
	}, nil
}
