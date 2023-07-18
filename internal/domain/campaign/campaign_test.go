package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "xbolinha"
	content  = "Body"
	contacts = []string{"email1@gmail.com", "email2@gmail.com", "email3@gmail.com"}
)

func TestCampaign_createCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Id, "ciqvh52navij03c99p5g")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func TestNewCampaignIdNotNull(t *testing.T) {
	assert := assert.New(t)
	campaign, e := NewCampaign(name, content, contacts)

	if e == "" {
		print(e)
	}
	assert.NotNil(campaign.Id)
}

func TestNewCampaign_CreateOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)
	campaign, e := NewCampaign(name, content, contacts)
	if e != "" {
		print(e)
	}

	assert.Greater(campaign.CreateOn, now)
}

func TestNewCampaign_MustValidadeName(t *testing.T) {
	assert := assert.New(t)

	e, err := NewCampaign("", content, contacts)

	if e != nil {
		print(e)
	}

	assert.Equal("variable is null", err)
}

func TestNewCampaign_MustValidadeContent(t *testing.T) {
	assert := assert.New(t)

	e, err := NewCampaign(name, "", contacts)

	if e != nil {
		print(e)
	}

	assert.Equal("variable is null", err)
}

func TestNewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	e, err := NewCampaign(name, content, []string{})

	if e != nil {
		print(e)
	}
	assert.Equal("variable is null", err)
}
