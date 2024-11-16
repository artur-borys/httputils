package tripperwares

import "net/http"

// Tripperware wraps a [http.RoundTripper]
type Tripperware func(rt http.RoundTripper) RoundTripperFunction

// Chain allows to chain multiple [Tripperware].
//
// If rt is nil, clone of [http.DefaultTransport] will be used as the base roundtripper
func Chain(rt http.RoundTripper, wares ...Tripperware) http.RoundTripper {
	if rt == nil { // coverage-ignore
		rt = http.DefaultTransport.(*http.Transport).Clone()
	}

	for _, w := range wares {
		rt = w(rt)
	}

	return rt
}
