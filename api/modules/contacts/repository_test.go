package contacts

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_repository_get_by_id(t *testing.T) {
	idExpected := "xer4"
	createdAtExpected := time.Now()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "active", "created_at", "user_id"}).
		AddRow(idExpected, "test", "test@test.com.br", true, createdAtExpected, "2sd2")
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select "id", "name", "email", "active", "created_at", "user_id" from users`).
		WithArgs(idExpected).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	contact, _ := repository.GetBy(idExpected)

	assert.Equal(t, idExpected, contact.Id)
	assert.Equal(t, "test", contact.Name)
	assert.Equal(t, "test@test.com.br", contact.Email)
	assert.Equal(t, true, contact.Active)
	assert.Equal(t, createdAtExpected, contact.CreatedAt)
	assert.Equal(t, "2sd2", contact.UserId)
}

func Test_repository_list_by_id(t *testing.T) {
	userIdExpected := "xer4"
	createdAtExpected := time.Now()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "active", "created_at", "user_id"}).
		AddRow(userIdExpected, "test", "test@test.com.br", true, createdAtExpected, "2sd2")
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select "id", "name", "email", "active", "created_at", "user_id" from users`).
		WithArgs(userIdExpected).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	contacts, _ := repository.ListBy(userIdExpected)

	assert.Equal(t, userIdExpected, contacts[0].Id)
	assert.Equal(t, "test", contacts[0].Name)
	assert.Equal(t, "test@test.com.br", contacts[0].Email)
	assert.Equal(t, true, contacts[0].Active)
	assert.Equal(t, createdAtExpected, contacts[0].CreatedAt)
	assert.Equal(t, "2sd2", contacts[0].UserId)
}

func Test_repository_create(t *testing.T) {
	contact := Contact{
		Id:        "xpt1",
		Name:      "Test",
		Email:     "test@test.com",
		Active:    true,
		CreatedAt: time.Now(),
		UserId:    "ee112",
	}
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("INSERT INTO contacts").
		ExpectExec().
		WithArgs(contact.Id, contact.Name, contact.Email, contact.Active, contact.CreatedAt, contact.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Create(contact)

	assert.Nil(t, err)
}

func Test_repository_save(t *testing.T) {
	contact := Contact{
		Id:        "xpt1",
		Name:      "Test",
		Email:     "test@test.com",
		Active:    false,
		CreatedAt: time.Now(),
		UserId:    "324d",
	}
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("UPDATE contacts").
		ExpectExec().
		WithArgs(contact.Name, contact.Active, contact.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Save(contact)

	assert.Nil(t, err)
}
