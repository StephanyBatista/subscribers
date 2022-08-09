package handlers_test

import (
	"net/http"
	"subscribers/domain/clients"
	"subscribers/domain/users"
	"subscribers/helpers"
	"subscribers/helpers/fake"
	"subscribers/web/handlers"
	"testing"

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

func createNewClient(name string, email string, userId string) clients.Client {
	entity := clients.NewClient(name, email, userId)
	fake.DB.Create(&entity)
	return *entity
}

func Test_client_post_validate_token(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/clients", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_client_post_validate_fields(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/clients", nil, fake.GenerateAnyToken())

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Email' is required")
}

func Test_client_post_save_new_client(t *testing.T) {
	fake.Build()
	body := handlers.ClientRequest{
		Name:  "Teste",
		Email: "teste@teste.com.br",
	}

	w := fake.MakeTestHTTP("POST", "/clients", body, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_client_post_show_error_when_not_create(t *testing.T) {
	fake.Build()
	mock := &fake.RepositoryMock[clients.Client]{}
	mock.ReturnsCreate = false
	fake.DI.ClientHandler.ClientRepository = mock
	body := handlers.ClientRequest{
		Name:  "Teste",
		Email: "teste@teste.com.br",
	}

	w := fake.MakeTestHTTP("POST", "/clients", body, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_client_get_all_validate_token(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("GET", "/clients", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_client_get_all_clients(t *testing.T) {
	fake.Build()
	user := createNewUser()
	createNewClient("teste2", "test1@teste.com.br", user.ID)
	createNewClient("teste2", "test2@teste.com.br", user.ID)
	amountOfClients := 2

	w := fake.MakeTestHTTP("GET", "/clients", nil, fake.GenerateTokenWithUser(user))

	clientsOfUser := helpers.BufferToObj[[]clients.Client](w.Body)
	assert.Equal(t, amountOfClients, len(clientsOfUser))
}

func Test_client_get_all_not_return_clients_from_others_users(t *testing.T) {
	fake.Build()
	currentUser := createNewUser()
	createNewClient("teste1", "test1@teste.com.br", currentUser.ID)
	createNewClient("teste2", "test2@teste.com.br", currentUser.ID)
	anotherUser := createNewUser()
	createNewClient("teste3", "test3@teste.com.br", anotherUser.ID)
	amountOfClients := 2

	w := fake.MakeTestHTTP("GET", "/clients", nil, fake.GenerateTokenWithUser(currentUser))

	clientsOfUser := helpers.BufferToObj[[]clients.Client](w.Body)
	assert.Equal(t, amountOfClients, len(clientsOfUser))
}

func Test_client_get_by_id_validate_token(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("GET", "/clients/xpter", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_client_get_by_id_return_client(t *testing.T) {
	fake.Build()
	currentUser := createNewUser()
	entity := createNewClient("teste2", "test2@teste.com.br", currentUser.ID)

	w := fake.MakeTestHTTP("GET", "/clients/"+entity.ID, nil, fake.GenerateTokenWithUser(currentUser))

	response := helpers.BufferToObj[clients.Client](w.Body)
	assert.Contains(t, response.ID, entity.ID)
	assert.Contains(t, response.Name, entity.Name)
	assert.Contains(t, response.Email, entity.Email)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_client_get_by_id_not_return_clients_from_others_users(t *testing.T) {
	fake.Build()
	currentUser := createNewUser()
	entity := createNewClient("teste2", "test2@teste.com.br", currentUser.ID)
	fake.DB.Create(&entity)

	w := fake.MakeTestHTTP("GET", "/clients/"+entity.ID, nil, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusNotFound, w.Code)
}
