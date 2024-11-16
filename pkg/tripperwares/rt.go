package tripperwares

import "net/http"

// RoundTripperFunction is similar to [http.HandlerFunc]
//
// Remember to adhere to [http.RoundTripper] interface requirements
// (i.e do not modify the original request - create a copy instead)
type RoundTripperFunction func(req *http.Request) (*http.Response, error)

// RoundTrip calls the [RoundTripperFunction] to implement [http.RoundTripper] interface
func (rtf RoundTripperFunction) RoundTrip(req *http.Request) (*http.Response, error) {
	return rtf(req)
}
