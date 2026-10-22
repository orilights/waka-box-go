package wakabox

import "net/http"

// wakatimeOAuthTransport replaces the legacy client's Basic authentication.
type wakatimeOAuthTransport struct {
	token string
	base  http.RoundTripper
}

func (t *wakatimeOAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	request := req.Clone(req.Context())
	request.Header.Set("Authorization", "Bearer "+t.token)
	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}
	return base.RoundTrip(request)
}
