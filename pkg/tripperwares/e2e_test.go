package tripperwares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	expectedUsername, expectedPassword := "username", "password"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actualUsername, actualPassword, _ := r.BasicAuth()

		assert.Equal(t, expectedUsername, actualUsername)
		assert.Equal(t, expectedPassword, actualPassword)

		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	client := server.Client()

	trChain := Chain(client.Transport, NewBasicAuthTW(expectedUsername, expectedPassword))
	client.Transport = trChain

	resp, err := client.Get(server.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
