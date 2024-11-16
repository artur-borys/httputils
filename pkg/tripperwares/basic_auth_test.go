package tripperwares

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuthTW(t *testing.T) {
	t.Run("adds username and password to the request", func(t *testing.T) {
		tw := NewBasicAuthTW("username", "password")

		baseRt := RoundTripperFunction(func(req *http.Request) (*http.Response, error) {
			if req.Body != nil {
				defer req.Body.Close()
			}

			return &http.Response{StatusCode: 200, ContentLength: 0, Request: req}, nil
		})

		rt := tw(baseRt)

		req, err := http.NewRequest("GET", "https://example.com", nil)
		origReq := *req

		assert.NoError(t, err)

		resp, err := rt.RoundTrip(req)

		assert.NoError(t, err)
		// original request left unmodified
		assert.Equal(t, origReq, *req)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, int64(0), resp.ContentLength)

		actualUsername, actualPassword, _ := resp.Request.BasicAuth()

		assert.Equal(t, "username", actualUsername)
		assert.Equal(t, "password", actualPassword)
	})
}
