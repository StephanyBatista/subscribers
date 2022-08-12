package handlers_test

import (
	"net/http"
	"subscribers/domain/contacts"
	"subscribers/domain/users"
	"subscribers/helpers"
	"subscribers/helpers/fake"
	"subscribers/web/handlers"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func createNewUser() users.User {
	var user *users.User
	result := fake.DB.Where(users.User{Email: "xaspqe@teste.com.br"}).Find(user)
	if result.RowsAffected == 0 {
		user, _ =
			users.NewUser("Teste", "xaspqe@teste.com.br", "password123")
		fake.DB.Create(user)
	}

	return *user
}

func createNewContact(userId string) contacts.Contact {
	entity := contacts.NewContact(gofakeit.Name(), gofakeit.Email(), userId)
	fake.DB.Create(&entity)
	return *entity
}

func Test_contact_post_validate_token(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/contacts", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_contact_post_validate_fields(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/contacts", nil, fake.GenerateAnyToken())

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Email' is required")
}

func Test_contact_post_save_new_client(t *testing.T) {
	fake.Build()
	body := handlers.ContactRequest{
		Name:  "Teste",
		Email: "teste@teste.com.br",
	}

	w := fake.MakeTestHTTP("POST", "/contacts", body, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_contact_post_show_error_when_not_create(t *testing.T) {
	fake.Build()
	mock := &fake.RepositoryMock[contacts.Contact]{}
	mock.ReturnsCreate = false
	fake.DI.ContactHandler.ContactRepository = mock
	body := handlers.ContactRequest{
		Name:  "Teste",
		Email: "teste@teste.com.br",
	}

	w := fake.MakeTestHTTP("POST", "/contacts", body, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_contact_get_all_validate_token(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("GET", "/contacts", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_contact_get_all_clients(t *testing.T) {
	fake.Build()
	user := createNewUser()
	createNewContact(user.ID)
	createNewContact(user.ID)
	amountOfClients := 2

	w := fake.MakeTestHTTP("GET", "/contacts", nil, fake.GenerateTokenWithUser(user))

	clientsOfUser := helpers.BufferToObj[[]contacts.Contact](w.Body)
	assert.Equal(t, amountOfClients, len(clientsOfUser))
}

func Test_contact_get_all_not_return_clients_from_others_users(t *testing.T) {
	fake.Build()
	currentUser := createNewUser()
	createNewContact(currentUser.ID)
	createNewContact(currentUser.ID)
	anotherUser := createNewUser()
	createNewContact(anotherUser.ID)
	amountOfClients := 2

	w := fake.MakeTestHTTP("GET", "/contacts", nil, fake.GenerateTokenWithUser(currentUser))

	clientsOfUser := helpers.BufferToObj[[]contacts.Contact](w.Body)
	assert.Equal(t, amountOfClients, len(clientsOfUser))
}

func Test_contact_get_by_id_validate_token(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("GET", "/contacts/xpter", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_contact_get_by_id_return_client(t *testing.T) {
	fake.Build()
	currentUser := createNewUser()
	entity := createNewContact(currentUser.ID)

	w := fake.MakeTestHTTP("GET", "/contacts/"+entity.ID, nil, fake.GenerateTokenWithUser(currentUser))

	response := helpers.BufferToObj[contacts.Contact](w.Body)
	assert.Contains(t, response.ID, entity.ID)
	assert.Contains(t, response.Name, entity.Name)
	assert.Contains(t, response.Email, entity.Email)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_contact_get_by_id_not_return_contact_from_others_users(t *testing.T) {
	fake.Build()
	currentUser := createNewUser()
	entity := createNewContact(currentUser.ID)
	fake.DB.Create(&entity)

	w := fake.MakeTestHTTP("GET", "/contacts/"+entity.ID, nil, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusNotFound, w.Code)
}
