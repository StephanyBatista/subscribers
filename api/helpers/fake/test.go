package fake

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"subscribers/helpers"
	"subscribers/web/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FakeEnvs() {
	os.Setenv("sub_database", "sqlite")
	os.Setenv("sub_salt_hash", "6")
}

var DI *helpers.DI = nil
var DB *gorm.DB = nil

func Build() {

	FakeEnvs()
	DI = helpers.NewDI()
	DB = DI.DB
}

func CreateHTTPTest(method string, route string, obj interface{}, token string) *httptest.ResponseRecorder {

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
