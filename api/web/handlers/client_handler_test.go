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

func Test_Client_Post_Validate_Parameter_UserId(t *testing.T) {
	fake.Build()
	body := handlers.ClientRequest{
		Name:  "Teste",
		Email: "teste@teste.com.br",
	}

	w := fake.MakeTestHTTP("POST", "/clients/user_invalid", body, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "User not found")
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_Client_Post_Validate_Fields(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/clients/"+createNewUser().ID, nil, "")

	response := helpers.BufferToString(w.Body)
	assert.Contains(t, response, "'Name' is required")
	assert.Contains(t, response, "'Email' is required")
}

func Test_Client_Post_Save_New_Client_Using_Param_UserId(t *testing.T) {
	fake.Build()
	body := handlers.ClientRequest{
		Name:  "Teste",
		Email: "teste@teste.com.br",
	}

	w := fake.MakeTestHTTP("POST", "/clients/"+createNewUser().ID, body, "")

	assert.Equal(t, http.StatusCreated, w.Code)
}

func Test_Client_Get_All_Clients(t *testing.T) {
	fake.Build()
	user := createNewUser()
	createNewClient("teste2", "test1@teste.com.br", user.ID)
	createNewClient("teste2", "test2@teste.com.br", user.ID)
	amountOfClients := 2

	w := fake.MakeTestHTTP("GET", "/clients", nil, fake.GenerateTokenWithUser(user))

	clientsOfUser := helpers.BufferToObj[[]clients.Client](w.Body)
	assert.Equal(t, amountOfClients, len(clientsOfUser))
}

func Test_Client_Get_All_Not_Return_Clients_From_Others_Users(t *testing.T) {
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

func Test_Client_Get_By_Id_Return_Client(t *testing.T) {
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

func Test_Client_Get_By_Id_Not_Return_Clients_From_Others_Users(t *testing.T) {
	fake.Build()
	currentUser := createNewUser()
	entity := createNewClient("teste2", "test2@teste.com.br", currentUser.ID)
	fake.DB.Create(&entity)

	w := fake.MakeTestHTTP("GET", "/clients/"+entity.ID, nil, fake.GenerateAnyToken())

	assert.Equal(t, http.StatusNotFound, w.Code)
}
