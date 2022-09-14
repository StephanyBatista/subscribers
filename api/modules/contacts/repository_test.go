package contacts

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var contactExpected Contact = Contact{
	Id:        "xpt1",
	Name:      "Test",
	Email:     "test@test.com",
	Active:    true,
	CreatedAt: time.Now(),
	UserId:    "ee112",
}

func Test_repository_get_by_id(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "active", "created_at", "user_id"}).
		AddRow(contactExpected.Id, contactExpected.Name, contactExpected.Email, contactExpected.Active, contactExpected.CreatedAt, contactExpected.UserId)
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select "id", "name", "email", "active", "created_at", "user_id" from users`).
		WithArgs(contactExpected.Id).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	contact, _ := repository.GetBy(contactExpected.Id)

	assert.Equal(t, contactExpected.Id, contact.Id)
	assert.Equal(t, contactExpected.Name, contact.Name)
	assert.Equal(t, contactExpected.Email, contact.Email)
	assert.Equal(t, contactExpected.Active, contact.Active)
	assert.Equal(t, contactExpected.CreatedAt, contact.CreatedAt)
	assert.Equal(t, contactExpected.UserId, contact.UserId)
}

func Test_repository_list_by_id(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "active", "created_at", "user_id"}).
		AddRow(contactExpected.Id, contactExpected.Name, contactExpected.Email, contactExpected.Active, contactExpected.CreatedAt, contactExpected.UserId)
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery(`select "id", "name", "email", "active", "created_at", "user_id" from users`).
		WithArgs(contactExpected.UserId).
		WillReturnRows(rows)
	repository := Repository{DB: db}

	contacts, _ := repository.ListBy(contactExpected.UserId)

	assert.Equal(t, contactExpected.Id, contacts[0].Id)
	assert.Equal(t, contactExpected.Name, contacts[0].Name)
	assert.Equal(t, contactExpected.Email, contacts[0].Email)
	assert.Equal(t, contactExpected.Active, contacts[0].Active)
	assert.Equal(t, contactExpected.CreatedAt, contacts[0].CreatedAt)
	assert.Equal(t, contactExpected.UserId, contacts[0].UserId)
}

func Test_repository_create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("INSERT INTO contacts").
		ExpectExec().
		WithArgs(contactExpected.Id, contactExpected.Name, contactExpected.Email, contactExpected.Active, contactExpected.CreatedAt, contactExpected.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Create(contactExpected)

	assert.Nil(t, err)
}

func Test_repository_save(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mock.
		ExpectPrepare("UPDATE contacts").
		ExpectExec().
		WithArgs(contactExpected.Name, contactExpected.Active, contactExpected.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := Repository{DB: db}

	err := repository.Save(contactExpected)

	assert.Nil(t, err)
}
