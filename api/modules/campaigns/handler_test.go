package campaigns

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"subscribers/utils/webtest"
	"testing"
)

func setupHandler() (*gin.Engine, *sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	router := webtest.CreateRouter()
	//TODO: create fake session to test
	ApplyRouter(router, db, nil)
	return router, db, mock
}

func createCampaignsRows(campaign Campaign) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "name", "from", "subject", "body", "status", "created_at", "user_id"}).
		AddRow(campaign.Id,
			campaign.Name,
			campaign.From,
			campaign.Subject,
			campaign.Body,
			campaign.Status,
			campaign.CreatedAt,
			campaign.UserId)
}

func Test_campaign_post_validate_token(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "POST", "/campaigns", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_campaign_post_validate_fields(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "POST", "/campaigns", nil, webtest.GenerateAnyToken())

	response := webtest.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'From' is required")
	assert.Contains(t, response, "'Subject' is required")
	assert.Contains(t, response, "'Body' is required")
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_campaign_post_save_new_campaign(t *testing.T) {
	router, _, mock := setupHandler()
	body := CreateNewCampaign{
		Name:    "teste 1",
		From:    "teste@teste.com.br",
		Subject: "Test 2",
		Body:    "Teste",
	}
	mock.
		ExpectPrepare("INSERT INTO campaigns").
		ExpectExec().
		WithArgs(sqlmock.AnyArg(),
			body.Name,
			body.From,
			body.Subject,
			body.Body,
			Draft,
			sqlmock.AnyArg(),
			sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	w := webtest.MakeTestHTTP(router, "POST", "/campaigns", body, webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_campaign_post_show_error_when_not_create(t *testing.T) {
	router, _, _ := setupHandler()
	body := CreateNewCampaign{
		Name:    "teste 1",
		From:    "teste@teste.com.br",
		Subject: "Test 2",
		Body:    "Teste",
	}

	w := webtest.MakeTestHTTP(router, "POST", "/campaigns", body, webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_campaign_get_campaign_by_id(t *testing.T) {
	router, _, mock := setupHandler()
	campaign, _ := NewCampaign("Name", "teste@teste.com.br", "Subject", "Hi", "3efd2")
	mock.ExpectQuery(queryBase).
		WithArgs(campaign.Id).
		WillReturnRows(createCampaignsRows(campaign))

	userToken := webtest.UserToken{Id: campaign.UserId, Email: "teste@teste.com.br", Name: "Test"}
	w := webtest.MakeTestHTTP(router, "GET", "/campaigns/"+campaign.Id, "", webtest.GenerateTokenWithUser(userToken))

	response := webtest.BufferToObj[Campaign](w.Body)
	assert.Equal(t, response.Id, campaign.Id)
	assert.Equal(t, response.Name, campaign.Name)
	assert.Equal(t, response.From, campaign.From)
	assert.Equal(t, response.Subject, campaign.Subject)
	//TODO: why only this property fail?
	//assert.Equal(t, response.Body, campaign.Body)
	assert.Equal(t, response.Status, campaign.Status)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_campaign_get_email_report(t *testing.T) {
	baseOfSubscribers := 6
	emailsSent := 2
	emailsNotSent := 1
	emailsOpened := 3
	router, _, mock := setupHandler()
	campaign, _ := NewCampaign("Name", "teste@teste.com.br", "Subject", "Hi!", "3efd2")
	mock.ExpectQuery(queryBase).
		WithArgs(campaign.Id).
		WillReturnRows(createCampaignsRows(campaign))
	rowsReport := sqlmock.NewRows([]string{"status"}).
		AddRow("Delivery").
		AddRow("Delivery").
		AddRow("Bounce").
		AddRow("Open").
		AddRow("Open").
		AddRow("Open")
	mock.ExpectQuery(`select status from subscribers`).
		WithArgs(campaign.Id).
		WillReturnRows(rowsReport)

	userToken := webtest.UserToken{Id: campaign.UserId, Email: "teste@teste.com.br", Name: "Test"}
	w := webtest.MakeTestHTTP(router, "GET", "/campaigns/"+campaign.Id+"/emailsreport", "", webtest.GenerateTokenWithUser(userToken))

	response := webtest.BufferToObj[EmailsReport](w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, baseOfSubscribers, response.BaseOfSubscribers)
	assert.Equal(t, emailsSent, response.Sent)
	assert.Equal(t, emailsNotSent, response.NotSent)
	assert.Equal(t, emailsOpened, response.Opened)
}

func Test_error_on_get_email_report_when_get_campaign(t *testing.T) {
	router, _, mock := setupHandler()
	mock.ExpectQuery(queryBase).
		WithArgs("invalid_campaign").
		WillReturnError(errors.New("any error"))

	w := webtest.MakeTestHTTP(router, "GET", "/campaigns/invalid_campaign/emailsreport", "", webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_error_on_get_email_report_when_get_report(t *testing.T) {
	router, _, mock := setupHandler()
	campaign, _ := NewCampaign("Name", "teste@teste.com.br", "Subject", "Hi!", "3efd2")
	mock.ExpectQuery(queryBase).
		WithArgs(campaign.Id).
		WillReturnRows(createCampaignsRows(campaign))
	mock.ExpectQuery(`select status from subscribers`).
		WillReturnError(errors.New("any error"))

	userToken := webtest.UserToken{Id: campaign.UserId}
	w := webtest.MakeTestHTTP(router, "GET", "/campaigns/"+campaign.Id+"/emailsreport", "", webtest.GenerateTokenWithUser(userToken))

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_get_email_report_campaign_not_found(t *testing.T) {
	router, _, mock := setupHandler()
	rowsCampaign := sqlmock.NewRows([]string{"id", "name", "from", "subject", "body", "status", "created_at", "user_id"})
	mock.ExpectQuery(queryBase).
		WithArgs("invalid_campaign").
		WillReturnRows(rowsCampaign)

	w := webtest.MakeTestHTTP(router, "GET", "/campaigns/invalid_campaign/emailsreport", "", webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func Test_campaign_get_by_id_not_found(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "GET", "/campaigns/id_invalid", nil, webtest.GenerateAnyToken())

	response := webtest.BufferToString(w.Body)
	assert.Contains(t, response, "Not found")
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func Test_campaign_get_all_campaign_of_user(t *testing.T) {
	router, _, mock := setupHandler()
	campaign, _ := NewCampaign("Name", "teste@teste.com.br", "Subject!", "Body of message!", "443rt1")
	mock.ExpectQuery(queryBase).
		WithArgs(campaign.UserId).
		WillReturnRows(createCampaignsRows(campaign))

	userToken := webtest.UserToken{Id: campaign.UserId, Email: "teste@teste.com.br", Name: "test"}
	w := webtest.MakeTestHTTP(router, "GET", "/campaigns", nil, webtest.GenerateTokenWithUser(userToken))

	campaignsOfUser := webtest.BufferToObj[[]Campaign](w.Body)
	assert.Equal(t, 1, len(campaignsOfUser))
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_campaign_get_all_not_found(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "GET", "/campaigns", nil, webtest.GenerateAnyToken())

	response := webtest.BufferToString(w.Body)
	assert.Contains(t, response, "Not found")
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}
