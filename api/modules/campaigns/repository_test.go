package campaigns

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var campaignExpected Campaign = Campaign{
	Id:        "xpt1",
	Name:      "Test",
	From:      "test@test.com",
	Subject:   "subject",
	Body:      "Body",
	Status:    "Status",
	CreatedAt: time.Now(),
	UserId:    "ee112",
}

func Test_repository_get_by_id(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "name", "from", "subject", "body", "status", "created_at", "user_id"}).
		AddRow(campaignExpected.Id,
			campaignExpected.Name,
			campaignExpected.From,
			campaignExpected.Subject,
			campaignExpected.Body,
			campaignExpected.Status,
			campaignExpected.CreatedAt,
			campaignExpected.UserId)
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select "id", "name", "from", "subject", "body", "status", "created_at", "user_id" from campaigns`).
		WithArgs(campaignExpected.Id).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	campaign, _ := repository.GetBy(campaignExpected.Id)

	assert.Equal(t, campaignExpected.Id, campaign.Id)
	assert.Equal(t, campaignExpected.Name, campaign.Name)
	assert.Equal(t, campaignExpected.From, campaign.From)
	assert.Equal(t, campaignExpected.Subject, campaign.Subject)
	assert.Equal(t, campaignExpected.Body, campaign.Body)
	assert.Equal(t, campaignExpected.Status, campaign.Status)
	assert.Equal(t, campaignExpected.CreatedAt, campaign.CreatedAt)
	assert.Equal(t, campaignExpected.UserId, campaign.UserId)
}

func Test_repository_get_emails_report(t *testing.T) {
	baseOfSubscribers := 6
	emailsSent := 2
	emailsNotSent := 1
	emailsOpened := 3
	campaignId := "xt34"
	rows := sqlmock.NewRows([]string{"status"}).
		AddRow("Delivery").
		AddRow("Delivery").
		AddRow("Bounce").
		AddRow("Open").
		AddRow("Open").
		AddRow("Open")
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select status from subscribers`).
		WithArgs(campaignId).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	emailsReport, _ := repository.GetEmailsReport(campaignId)

	assert.Equal(t, baseOfSubscribers, emailsReport.BaseOfSubscribers)
	assert.Equal(t, emailsSent, emailsReport.Sent)
	assert.Equal(t, emailsNotSent, emailsReport.NotSent)
	assert.Equal(t, emailsOpened, emailsReport.Opened)
}

func Test_repository_list_by_id(t *testing.T) {

	rows := sqlmock.NewRows([]string{"id", "name", "from", "subject", "body", "status", "created_at", "user_id"}).
		AddRow(campaignExpected.Id,
			campaignExpected.Name,
			campaignExpected.From,
			campaignExpected.Subject,
			campaignExpected.Body,
			campaignExpected.Status,
			campaignExpected.CreatedAt,
			campaignExpected.UserId)
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select "id", "name", "from", "subject", "body", "status", "created_at", "user_id" from campaigns`).
		WithArgs(campaignExpected.UserId).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	campaigns, _ := repository.ListBy(campaignExpected.UserId)

	assert.Equal(t, campaignExpected.Id, campaigns[0].Id)
	assert.Equal(t, campaignExpected.Name, campaigns[0].Name)
	assert.Equal(t, campaignExpected.From, campaigns[0].From)
	assert.Equal(t, campaignExpected.Subject, campaigns[0].Subject)
	assert.Equal(t, campaignExpected.Body, campaigns[0].Body)
	assert.Equal(t, campaignExpected.Status, campaigns[0].Status)
	assert.Equal(t, campaignExpected.CreatedAt, campaigns[0].CreatedAt)
	assert.Equal(t, campaignExpected.UserId, campaigns[0].UserId)
}

func Test_repository_create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("INSERT INTO campaigns").
		ExpectExec().
		WithArgs(campaignExpected.Id,
			campaignExpected.Name,
			campaignExpected.From,
			campaignExpected.Subject,
			campaignExpected.Body,
			campaignExpected.Status,
			campaignExpected.CreatedAt,
			campaignExpected.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Create(campaignExpected)

	assert.Nil(t, err)
}

func Test_repository_save(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("UPDATE campaigns").
		ExpectExec().
		WithArgs(campaignExpected.Name, campaignExpected.From, campaignExpected.Subject, campaignExpected.Body, campaignExpected.Status, campaignExpected.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Save(campaignExpected)

	assert.Nil(t, err)
}
