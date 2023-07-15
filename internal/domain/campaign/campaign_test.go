package campaign

import "testing"

func TestCampaign(t *testing.T){
	name := "xbolinha"
	content := "Body"
	contacts := []string{"email1@gmail.com", "email2@gmail.com", "email3@gmail.com"}
	campaign := NewCampaign(name, content, contacts)

	if campaign.Id != "1"
	t.Errorf("expected 1", campaign.Id)
}