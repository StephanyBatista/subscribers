package users

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_repository_get_by_email(t *testing.T) {
	emailExpected := "teste@teste.com"
	createdAtExpected := time.Now()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password_hash", "createdAt"}).
		AddRow("xpt1", "test", emailExpected, "serwswqer", createdAtExpected)
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select "id", "name", "email", "password_hash", "created_at" from users`).
		WithArgs(emailExpected).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	user, _ := repository.GetByEmail(emailExpected)

	assert.Equal(t, "xpt1", user.Id)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, emailExpected, user.Email)
	assert.Equal(t, createdAtExpected, user.CreatedAt)
}

func Test_repository_create(t *testing.T) {
	user := User{
		Id:           "xpt1",
		Name:         "Test",
		Email:        "test@test.com",
		PasswordHash: "xererdfkerwe3434df324",
		CreatedAt:    time.Now(),
	}
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("INSERT INTO users").
		ExpectExec().
		WithArgs(user.Id, user.Name, user.Email, user.PasswordHash, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Create(user)

	assert.Nil(t, err)
}

func Test_repository_save(t *testing.T) {
	user := User{
		Id:           "xpt1",
		Name:         "Test",
		Email:        "test@test.com",
		PasswordHash: "xererdfkerwe3434df324",
		CreatedAt:    time.Now(),
	}
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("UPDATE users").
		ExpectExec().
		WithArgs(user.Name, user.PasswordHash, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Save(user)

	assert.Nil(t, err)
}
