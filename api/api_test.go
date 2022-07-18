package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMissingBody(t *testing.T) {

	req, err := http.NewRequest("POST", "/url", nil)
	assert.Nil(t, err)
	response := executeRequest(req)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestMissingTarget(t *testing.T) {

	req, err := http.NewRequest("POST", "/url", bytes.NewReader([]byte{}))
	assert.Nil(t, err)
	response := executeRequest(req)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestInvalidTarget(t *testing.T) {
	b, err := json.Marshal(&struct {
		Target string `json:"target"`
	}{
		Target: "foo",
	})
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "/url", bytes.NewReader(b))
	assert.Nil(t, err)
	response := executeRequest(req)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	assert.Nil(t, err)
	response := executeRequest(req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	log := logrus.WithField("app", "test")
	New(log).ServeHTTP(rr, req)

	return rr
}
