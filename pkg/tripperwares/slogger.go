package tripperwares

import (
	"log/slog"
	"net/http"
)

// NewSloggerTW creates a Tripperware which logs both request and response metadata
//
// You must provide the base [slog.Logger] and the desired level ([slog.Level]).
//
// Do not use it if you care about performance!
func NewSloggerTW(baseLogger *slog.Logger, level slog.Level) Tripperware {
	return func(rt http.RoundTripper) RoundTripperFunction {
		return func(req *http.Request) (*http.Response, error) {
			logger := baseLogger.With("url", req.URL, "method", req.Method)

			logger.Log(req.Context(), level, "request", "content-length", req.ContentLength)

			// Not modifying req here, so it's fine to pass it directly
			resp, err := rt.RoundTrip(req)

			if err != nil {
				return resp, err
			}

			logger.Log(resp.Request.Context(), level, "response received", "status", resp.StatusCode, "content-length", resp.ContentLength)

			return resp, err
		}
	}
}
