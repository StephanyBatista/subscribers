package handlers_test

import (
	"fmt"
	"net/http"
	"subscribers/helpers"
	"subscribers/helpers/fake"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_file_post_validate_token(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/files", nil, "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_file_post_validate_fields(t *testing.T) {
	fake.Build()

	w := fake.MakeTestHTTP("POST", "/files", nil, fake.GenerateAnyToken())

	response := helpers.BufferToString(w.Body)
	fmt.Println(response)
	assert.Contains(t, response, "'keyId' is required")
	assert.Contains(t, response, "'kind' is required")
	assert.Contains(t, response, "'file' is required")
}

//TODO: Add test uploading file
