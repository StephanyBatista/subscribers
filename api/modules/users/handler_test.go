package users

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"net/http"
	"subscribers/modules/web"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupHandler() (*gin.Engine, *sql.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	router := web.CreateRouter()
	ApplyRouter(router, db)
	return router, db, mock
}

func Test_user_post_validate_fields_required(t *testing.T) {
	router, _, _ := setupHandler()

	w := web.MakeTestHTTP(router, "POST", "/users", nil, "")

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Email' is required")
	assert.Contains(t, response, "'Password' is required")
}

func Test_user_post_validate_invalid_email(t *testing.T) {
	router, _, _ := setupHandler()
	newUser := UserRequest{
		Name:     "Demo",
		Email:    "invalid",
		Password: "35 million",
	}

	w := web.MakeTestHTTP(router, "POST", "/users", newUser, "")

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "'Email' is invalid")
}

func Test_user_post_validate_when_email_is_being_used(t *testing.T) {
	router, _, mock := setupHandler()
	emailExpected := "teste@teste.com"
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password_hash", "createdAt"}).
		AddRow("xpt1", "test", emailExpected, "xptiwe", time.Now())
	mock.ExpectQuery(`select "id", "name", "email", "password_hash", "created_at" from users`).
		WithArgs(emailExpected).
		WillReturnRows(rows)

	newUser := UserRequest{
		Name:     "Demo",
		Email:    emailExpected,
		Password: "35 million",
	}

	w := web.MakeTestHTTP(router, "POST", "/users", newUser, "")

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "Email already saved")
}

func Test_user_post_save_new_user(t *testing.T) {
	router, _, mock := setupHandler()
	emailExpected := "teste1@teste.com"
	mock.ExpectPrepare("INSERT INTO users").
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(1, 1))

	newUser := UserRequest{
		Name:     "Demo",
		Email:    emailExpected,
		Password: "35 million",
	}

	w := web.MakeTestHTTP(router, "POST", "/users", newUser, "")

	assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
}

func Test_user_post_show_error_when_not_create(t *testing.T) {
	router, _, _ := setupHandler()
	newUser := UserRequest{
		Name:     "Demo",
		Email:    "teste1@teste.com",
		Password: "35 million",
	}

	w := web.MakeTestHTTP(router, "POST", "/users", newUser, "")

	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func Test_user_get_info(t *testing.T) {
	router, _, _ := setupHandler()
	userToken := web.UserToken{Name: "test1", Email: "test1@test.com"}

	w := web.MakeTestHTTP(router, "GET", "/users/info", nil, web.GenerateTokenWithUser(userToken))

	result := web.BufferToString(w.Body)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Contains(t, result, userToken.Email)
	assert.Contains(t, result, userToken.Name)
}

func Test_user_get_info_validate_jwt(t *testing.T) {
	router, _, _ := setupHandler()

	w := web.MakeTestHTTP(router, "GET", "/users/info", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
}

func Test_user_change_password_validate_parameters(t *testing.T) {
	router, _, _ := setupHandler()
	userToken := web.UserToken{Name: "test1", Email: "test1@test.com"}

	w := web.MakeTestHTTP(router, "PATCH", "/users/changepassword", nil, web.GenerateTokenWithUser(userToken))

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "'OldPassword' is required")
	assert.Contains(t, response, "'NewPassword' is required")
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func Test_user_change_password_check_old_password(t *testing.T) {
	router, _, mock := setupHandler()
	userToken := web.UserToken{Name: "test1", Email: "test1@test.com"}
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password_hash", "createdAt"}).
		AddRow("xpt1", "test", userToken.Email, "password_different", time.Now())
	mock.ExpectQuery(`select "id", "name", "email", "password_hash", "created_at" from users`).
		WithArgs(userToken.Email).
		WillReturnRows(rows)
	changePassword := UserChangePasswordRequest{NewPassword: "test", OldPassword: "password"}

	w := web.MakeTestHTTP(router, "PATCH", "/users/changepassword", changePassword, web.GenerateTokenWithUser(userToken))

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "old password invalid")
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func Test_user_must_change_password(t *testing.T) {
	router, _, mock := setupHandler()
	oldPassword := "password 2"
	userToken := web.UserToken{Name: "test1", Email: "test1@test.com"}
	user, _ := NewUser(userToken.Name, userToken.Email, oldPassword)
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password_hash", "createdAt"}).
		AddRow("xpt1", "test", userToken.Email, user.PasswordHash, time.Now())
	mock.ExpectQuery(`select "id", "name", "email", "password_hash", "created_at" from users`).
		WithArgs(userToken.Email).
		WillReturnRows(rows)
	changePassword := UserChangePasswordRequest{NewPassword: "test", OldPassword: oldPassword}

	w := web.MakeTestHTTP(router, "PATCH", "/users/changepassword", changePassword, web.GenerateTokenWithUser(userToken))

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func Test_token_post_validate_fields_required(t *testing.T) {
	router, _, _ := setupHandler()

	w := web.MakeTestHTTP(router, "POST", "/token", nil, "")

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "'Email' is required")
	assert.Contains(t, response, "'Password' is required")
}

func Test_token_post_user_not_found(t *testing.T) {
	router, _, _ := setupHandler()
	body := LoginRequest{
		Email:    "test1@teste.com.br",
		Password: "35 million",
	}

	w := web.MakeTestHTTP(router, "POST", "/token", body, "")

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "User not found")
	assert.Equal(t, http.StatusForbidden, w.Result().StatusCode)
}

func Test_token_post_generate_jwt(t *testing.T) {
	router, _, mock := setupHandler()
	passwordExpected := "35 million"
	user, _ := NewUser("Teste", "test1@teste.com.br", passwordExpected)
	body := LoginRequest{
		Email:    user.Email,
		Password: passwordExpected,
	}
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password_hash", "createdAt"}).
		AddRow("xpt1", "test", body.Email, user.PasswordHash, time.Now())
	mock.ExpectQuery(`select "id", "name", "email", "password_hash", "created_at" from users`).
		WithArgs(body.Email).
		WillReturnRows(rows)

	w := web.MakeTestHTTP(router, "POST", "/token", body, "")

	response := web.BufferToString(w.Body)
	assert.Contains(t, response, "token")
	assert.Contains(t, response, "expiresAt")
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
