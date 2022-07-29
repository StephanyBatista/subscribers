package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
)

func FakeEnvs() {
	os.Setenv("sub_database", "sqlite")
	os.Setenv("sub_salt_hash", "6")
}

func CreateHTTPTest(method string, route string, handlerFunc gin.HandlerFunc, obj interface{}) *httptest.ResponseRecorder {
	r := gin.Default()
	if method == "POST" {
		r.POST(route, handlerFunc)
	} else if method == "GET" {
		r.GET(route, handlerFunc)
	}

	var req *http.Request
	if obj != nil {
		jsonValue, _ := json.Marshal(obj)
		buffer := bytes.NewBuffer(jsonValue)
		req, _ = http.NewRequest(method, route, buffer)
	} else {
		req, _ = http.NewRequest(method, route, nil)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
