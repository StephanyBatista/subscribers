package fake

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"subscribers/domain/users"
	"subscribers/helpers"
	"subscribers/web/auth"
	"subscribers/web/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FakeEnvs() {
	os.Setenv("sub_database", "sqlite:memory")
	os.Setenv("sub_salt_hash", "6")
}

var DI *helpers.DI = nil
var DB *gorm.DB = nil

func Build() {

	FakeEnvs()
	DI = helpers.NewDI()
	DB = DI.DB
}

func GenerateAnyToken() string {
	token, _, _ := auth.GenerateJWT("xpto", "teste@teste.com.br", "test")
	return token
}

func GenerateTokenWithUserId(userId string) string {
	token, _, _ := auth.GenerateJWT(userId, "teste@teste.com.br", "test")
	return token
}

func GenerateTokenWithUser(user users.User) string {
	token, _, _ := auth.GenerateJWT(user.ID, user.Email, user.Name)
	return token
}

func MakeTestHTTP(method string, route string, obj interface{}, token string) *httptest.ResponseRecorder {

	if DI == nil || DB == nil {
		panic("Before call method you must call fake.Build")
	}

	gin.SetMode(gin.ReleaseMode)

	r := routers.CreateRouter(DI)

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
	r.ServeHTTP(w, req)
	return w
}
