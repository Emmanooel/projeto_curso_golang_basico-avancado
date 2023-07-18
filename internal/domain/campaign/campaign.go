package campaign

import (
	"time"

	"github.com/rs/xid"
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

func NewCampaign(name string, content string, emails []string) (*Campaign, string) {

	if name == "" || content == "" || len(emails) == 0 {
		return nil, "variable is null"
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		Id:       xid.New().String(),
		Name:     name,
		Content:  content,
		CreateOn: time.Now(),
		Contacts: contacts,
	}, ""
}
