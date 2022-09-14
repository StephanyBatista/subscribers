package campaigns

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"subscribers/modules/web"
	"testing"
	"time"
)

type AnyString struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func setupHandler() (*gin.Engine, *sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	router := web.CreateRouter()
	ApplyRouter(router, db)
	return router, db, mock
}

func Test_campaign_post_validate_token(t *testing.T) {
	router, _, _ := setupHandler()

	w := web.MakeTestHTTP(router, "POST", "/campaigns", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_campaign_post_validate_fields(t *testing.T) {
	router, _, _ := setupHandler()

	w := web.MakeTestHTTP(router, "POST", "/campaigns", nil, web.GenerateAnyToken())

	response := web.BufferToString(w.Body)
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
		WithArgs(AnyString{},
			body.Name,
			body.From,
			body.Subject,
			body.Body,
			Draft,
			AnyTime{},
			AnyString{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	w := web.MakeTestHTTP(router, "POST", "/campaigns", body, web.GenerateAnyToken())

	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_campaign_post_show_erro_when_not_create(t *testing.T) {
	router, _, _ := setupHandler()
	body := CreateNewCampaign{
		Name:    "teste 1",
		From:    "teste@teste.com.br",
		Subject: "Test 2",
		Body:    "Teste",
	}

	w := web.MakeTestHTTP(router, "POST", "/campaigns", body, web.GenerateAnyToken())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_campaign_get_campaign_by_id(t *testing.T) {
	router, _, mock := setupHandler()
	campaign, _ := NewCampaign("Name", "teste@teste.com.br", "Subject", "Hi!", "3efd2")
	rows := sqlmock.NewRows([]string{"id", "name", "from", "subject", "body", "status", "created_at", "user_id"}).
		AddRow(campaign.Id,
			campaign.Name,
			campaign.From,
			campaign.Subject,
			campaign.Body,
			campaign.Status,
			campaign.CreatedAt,
			campaign.UserId)
	mock.ExpectQuery(`select "id", "name", "from", "subject", "body", "status", "created_at", "user_id" from campaigns`).
		WithArgs(campaign.Id).
		WillReturnRows(rows)

	userToken := web.UserToken{Id: campaign.UserId, Email: "teste@teste.com.br", Name: "Test"}
	w := web.MakeTestHTTP(router, "GET", "/campaigns/"+campaign.Id, "", web.GenerateTokenWithUser(userToken))

	response := web.BufferToObj[CampaignResponse](w.Body)
	assert.Equal(t, response.ID, campaign.Id)
	assert.Equal(t, response.Name, campaign.Name)
	assert.Equal(t, response.From, campaign.From)
	assert.Equal(t, response.Subject, campaign.Subject)
	assert.Equal(t, response.Body, campaign.Body)
	assert.Equal(t, response.Status, campaign.Status)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_campaign_get_by_id_not_found(t *testing.T) {
	router, _, _ := setupHandler()

	w := web.MakeTestHTTP(router, "GET", "/campaigns/id_invalid", nil, web.GenerateAnyToken())

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "Not found")
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func Test_campaign_get_all_campaign_of_user(t *testing.T) {
	router, _, mock := setupHandler()
	campaign, _ := NewCampaign("Name", "teste@teste.com.br", "Subject!", "Body of message!", "443rt1")
	rows := sqlmock.NewRows([]string{"id", "name", "from", "subject", "body", "status", "created_at", "user_id"}).
		AddRow(campaign.Id,
			campaign.Name,
			campaign.From,
			campaign.Subject,
			campaign.Body,
			campaign.Status,
			campaign.CreatedAt,
			campaign.UserId)
	mock.ExpectQuery(`select "id", "name", "from", "subject", "body", "status", "created_at", "user_id" from campaigns`).
		WithArgs(campaign.UserId).
		WillReturnRows(rows)

	userToken := web.UserToken{Id: campaign.UserId, Email: "teste@teste.com.br", Name: "test"}
	w := web.MakeTestHTTP(router, "GET", "/campaigns", nil, web.GenerateTokenWithUser(userToken))

	campaignsOfUser := web.BufferToObj[[]Campaign](w.Body)
	assert.Equal(t, 1, len(campaignsOfUser))
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_campaign_get_all_not_found(t *testing.T) {
	router, _, _ := setupHandler()

	w := web.MakeTestHTTP(router, "GET", "/campaigns", nil, web.GenerateAnyToken())

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "Not found")
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}
