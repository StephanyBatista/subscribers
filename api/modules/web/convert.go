package web

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func BufferToString(buffer *bytes.Buffer) string {
	response, _ := ioutil.ReadAll(buffer)
	return string(response)
}

func BufferToObj[obj any](buffer *bytes.Buffer) obj {
	responseData, _ := ioutil.ReadAll(buffer)
	var result obj
	json.Unmarshal([]byte(responseData), &result)
	return result
}
