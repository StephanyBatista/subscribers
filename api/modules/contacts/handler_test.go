package contacts

import (
	"database/sql"
	"net/http"
	"subscribers/utils/webtest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupHandler() (*gin.Engine, *sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	router := webtest.CreateRouter()
	ApplyRouter(router, db)
	return router, db, mock
}

func Test_contact_post_validate_token(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "POST", "/contacts", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_contact_post_validate_fields(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "POST", "/contacts", nil, webtest.GenerateAnyToken())

	response := webtest.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Email' is required")
}

func Test_contact_post_save_new_client(t *testing.T) {
	router, _, mock := setupHandler()
	body := CreateNewContact{
		Name:  "Teste",
		Email: "teste@teste.com.br",
	}
	mock.ExpectPrepare("INSERT INTO contacts").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	w := webtest.MakeTestHTTP(router, "POST", "/contacts", body, webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_contact_post_show_error_when_not_create(t *testing.T) {
	router, _, _ := setupHandler()
	body := CreateNewContact{
		Name:  "Teste",
		Email: "teste@teste.com.br",
	}

	w := webtest.MakeTestHTTP(router, "POST", "/contacts", body, webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_contact_get_all_validate_token(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "GET", "/contacts", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_contact_get_all_clients(t *testing.T) {
	router, _, mock := setupHandler()
	amountOfClients := 1
	userToken := webtest.UserToken{Id: "xpt233", Email: "test@teste.com", Name: "test"}
	rows := sqlmock.NewRows([]string{"id", "name", "email", "active", "created_at", "user_id"}).
		AddRow("23s23", "test", "test@test.com.br", true, time.Now(), "2sd2")
	mock.ExpectQuery(`select "id", "name", "email", "active", "created_at", "user_id" from users`).
		WithArgs(userToken.Id).
		WillReturnRows(rows)

	w := webtest.MakeTestHTTP(router, "GET", "/contacts", nil, webtest.GenerateTokenWithUser(userToken))

	clientsOfUser := webtest.BufferToObj[[]Contact](w.Body)
	assert.Equal(t, amountOfClients, len(clientsOfUser))
}

func Test_contact_get_by_id_validate_token(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "GET", "/contacts/xpter", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_contact_get_by_id_return_client(t *testing.T) {
	router, _, mock := setupHandler()
	idExpected := "343"
	rows := sqlmock.NewRows([]string{"id", "name", "email", "active", "created_at", "user_id"}).
		AddRow(idExpected, "test", "test@test.com.br", true, time.Now(), "2sd2")
	mock.ExpectQuery(`select "id", "name", "email", "active", "created_at", "user_id" from users`).
		WithArgs(idExpected).
		WillReturnRows(rows)

	w := webtest.MakeTestHTTP(router, "GET", "/contacts/"+idExpected, nil, webtest.GenerateAnyToken())

	response := webtest.BufferToObj[Contact](w.Body)
	assert.Contains(t, response.Id, idExpected)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_contact_get_by_id_not_return_contact_from_others_users(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "GET", "/contacts/any_id", nil, webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func Test_contact_cancel_not_found(t *testing.T) {
	router, _, _ := setupHandler()

	w := webtest.MakeTestHTTP(router, "PATCH", "/contacts/any_id", nil, webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func Test_contact_cancel(t *testing.T) {
	router, _, mock := setupHandler()
	idExpected := "343"
	rows := sqlmock.NewRows([]string{"id", "name", "email", "active", "created_at", "user_id"}).
		AddRow(idExpected, "test", "test@test.com.br", true, time.Now(), "2sd2")
	mock.ExpectQuery(`select "id", "name", "email", "active", "created_at", "user_id" from users`).
		WithArgs(idExpected).
		WillReturnRows(rows)
	mock.
		ExpectPrepare("UPDATE contacts").
		ExpectExec().
		WithArgs("test", false, idExpected).
		WillReturnResult(sqlmock.NewResult(1, 1))

	w := webtest.MakeTestHTTP(router, "PATCH", "/contacts/"+idExpected+"/cancel", nil, webtest.GenerateAnyToken())

	assert.Equal(t, http.StatusOK, w.Code)
}
