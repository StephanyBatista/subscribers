package campaigns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_new_campaign(t *testing.T) {

	nameExpected := "test"
	fromExpected := "test@test.com.br"
	subjectExpected := "Campaign 1"
	bodyExpected := "Hi!"
	userId := "23121"
	campaign, _ := NewCampaign(nameExpected, fromExpected, subjectExpected, bodyExpected, userId)

	assert.Equal(t, nameExpected, campaign.Name)
	assert.Equal(t, fromExpected, campaign.From)
	assert.Equal(t, subjectExpected, campaign.Subject)
	assert.Equal(t, userId, campaign.UserId)
	assert.NotEmpty(t, campaign.Id)
	assert.NotEmpty(t, campaign.CreatedAt)
	assert.Equal(t, Draft, campaign.Status)
}

func Test_must_set_as_ready_when_call_ready(t *testing.T) {

	campaign := Campaign{}

	campaign.Ready()

	assert.Equal(t, Ready, campaign.Status)
}

func Test_must_validate_fields_when_create_new_campaign(t *testing.T) {

	_, errs := NewCampaign("", "", "", "", "")

	assert.Equal(t, "'Name' is required", errs[0].Error())
	assert.Equal(t, "'From' is required", errs[1].Error())
	assert.Equal(t, "'Subject' is required", errs[2].Error())
	assert.Equal(t, "'Body' is required", errs[3].Error())
	assert.Equal(t, "'UserId' is required", errs[4].Error())
}

func Test_must_validate_from_as_email_when_create_new_campaign(t *testing.T) {

	_, errs := NewCampaign("", "test", "", "", "")

	assert.Equal(t, "'From' invalid", errs[1].Error())
}
