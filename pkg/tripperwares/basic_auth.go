package tripperwares

import "net/http"

// NewBasicAuthTW returns a [Tripperware] which sets basic authentication header
func NewBasicAuthTW(username, password string) Tripperware {
	return func(rt http.RoundTripper) RoundTripperFunction {
		return func(req *http.Request) (*http.Response, error) {
			// Adhere RoundTripper interface - do not modify original request
			reqClone := req.Clone(req.Context())
			reqClone.SetBasicAuth(username, password)

			// Clone() does only shallow copy of the body, so it is assumed
			// the underlying http.Transport will close it
			return rt.RoundTrip(reqClone)
		}
	}
}
