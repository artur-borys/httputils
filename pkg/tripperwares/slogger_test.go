package tripperwares

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlogger(t *testing.T) {
	t.Run("Contains request and response metadata", func(t *testing.T) {
		outputBuffer := bytes.Buffer{}
		responseBody := []byte("Hello, World!")

		handler := slog.NewTextHandler(&outputBuffer, nil)
		logger := slog.New(handler)

		sloggerTW := NewSloggerTW(logger, slog.LevelInfo)

		rt := RoundTripperFunction(func(req *http.Request) (*http.Response, error) {
			if req.Body != nil {
				req.Body.Close()
			}

			return &http.Response{
				Request:       req,
				StatusCode:    200,
				ContentLength: int64(len(responseBody)),
				Body:          io.NopCloser(bytes.NewReader(responseBody)),
			}, nil
		})

		uut := sloggerTW(rt)

		req, err := http.NewRequest("GET", "https://example.com", &bytes.Buffer{})

		assert.NoError(t, err)

		resp, err := uut.RoundTrip(req)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		outputString := outputBuffer.String()

		assert.Contains(t, outputString, "INFO")
		assert.Contains(t, outputString, "method=GET")
		assert.Contains(t, outputString, "url=https://example.com")
		assert.Contains(t, outputString, "content-length=0")
		assert.Contains(t, outputString, "content-length="+fmt.Sprintf("%d", len(responseBody)))
		assert.Contains(t, outputString, "status=200")
	})

	t.Run("doesn't try to log resp metadata on error", func(t *testing.T) {
		outputBuffer := bytes.Buffer{}

		handler := slog.NewTextHandler(&outputBuffer, nil)
		logger := slog.New(handler)

		sloggerTW := NewSloggerTW(logger, slog.LevelInfo)

		rt := RoundTripperFunction(func(req *http.Request) (*http.Response, error) {
			if req.Body != nil {
				req.Body.Close()
			}

			return nil, errors.New("example error")
		})

		uut := sloggerTW(rt)

		req, err := http.NewRequest("GET", "https://example.com", &bytes.Buffer{})
		assert.NoError(t, err)

		resp, err := uut.RoundTrip(req)
		assert.Nil(t, resp)
		assert.Error(t, err)

		outputString := outputBuffer.String()
		assert.Contains(t, outputString, "INFO")
		assert.Contains(t, outputString, "method=GET")
		assert.Contains(t, outputString, "url=https://example.com")
		assert.Contains(t, outputString, "content-length=0")
		assert.NotContains(t, outputString, "status=")
	})
}
