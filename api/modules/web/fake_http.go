package web

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"subscribers/web/auth"
)

func FakeEnvs() {
	os.Setenv("sub_gotest", "true")
	os.Setenv("sub_salt_hash", "6")
}

func Build() {
	FakeEnvs()
}

type UserToken struct {
	Id    string
	Email string
	Name  string
}

func GenerateAnyToken() string {
	token, _, _ := auth.GenerateJWT("xpto", "teste@teste.com.br", "test")
	return token
}

func GenerateTokenWithUser(userToken UserToken) string {
	token, _, _ := auth.GenerateJWT(userToken.Id, userToken.Email, userToken.Name)
	return token
}

func MakeTestHTTP(server *gin.Engine, method string, route string, obj interface{}, token string) *httptest.ResponseRecorder {

	gin.SetMode(gin.ReleaseMode)

	var req *http.Request
	if obj != nil {
		jsonValue, _ := json.Marshal(obj)
		buffer := bytes.NewBuffer(jsonValue)
		req, _ = http.NewRequest(method, route, buffer)
	} else {
		req, _ = http.NewRequest(method, route, nil)
	}

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	return w
}
