package users

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var userExpected User = User{
	Id:           "xpt1",
	Name:         "Test",
	Email:        "test@test.com",
	PasswordHash: "xererdfkerwe3434df324",
	CreatedAt:    time.Now(),
}

func Test_repository_get_by_email(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password_hash", "createdAt"}).
		AddRow(userExpected.Id, userExpected.Name, userExpected.Email, userExpected.PasswordHash, userExpected.CreatedAt)
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select "id", "name", "email", "password_hash", "created_at" from users`).
		WithArgs(userExpected.Email).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	user, _ := repository.GetByEmail(userExpected.Email)

	assert.Equal(t, userExpected.Id, user.Id)
	assert.Equal(t, userExpected.Name, user.Name)
	assert.Equal(t, userExpected.Email, user.Email)
	assert.Equal(t, userExpected.CreatedAt, user.CreatedAt)
}

func Test_repository_create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("INSERT INTO users").
		ExpectExec().
		WithArgs(userExpected.Id, userExpected.Name, userExpected.Email, userExpected.PasswordHash, userExpected.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Create(userExpected)

	assert.Nil(t, err)
}

func Test_repository_save(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("UPDATE users").
		ExpectExec().
		WithArgs(userExpected.Name, userExpected.PasswordHash, userExpected.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Save(userExpected)

	assert.Nil(t, err)
}
