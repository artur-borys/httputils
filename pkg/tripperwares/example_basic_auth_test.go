package tripperwares_test

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/artur-borys/httputils/pkg/tripperwares"
)

func Example() {
	// Create a HTTP client instance with a chain of tripperwares as the Transport
	// If Chain has empty base RoundTripper, a clone of http.DefaultTransport will be used
	client := http.Client{
		Transport: tripperwares.Chain(nil, tripperwares.NewBasicAuthTW("somebody", "S3cure123!")),
	}

	resp, err := client.Get("https://example.com")

	if err != nil {
		slog.Error("Whoops, unexpected error", "error", err)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		slog.Error("Whoops, unexpected error", "error", err)
	}

	slog.Info(string(data))
}
