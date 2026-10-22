package wakabox

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type oauthTestTransport func(*http.Request) (*http.Response, error)

func (f oauthTestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestOAuthStatsRequest(t *testing.T) {
	previous := http.DefaultTransport
	defer func() { http.DefaultTransport = previous }()
	called := false
	http.DefaultTransport = oauthTestTransport(func(req *http.Request) (*http.Response, error) {
		called = true
		if got := req.Header.Get("Authorization"); got != "Bearer test-oauth-token" {
			t.Errorf("unexpected Authorization: %q", got)
		}
		if req.URL.Path != "/api/v1/users/current/stats/last_7_days" || req.URL.RawQuery != "" {
			t.Errorf("unexpected URL: %s", req.URL)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(strings.NewReader(`{"data":{"languages":[{"name":"Go","text":"1 hr","percent":100}]}}`)),
		}, nil
	})
	box := NewBox("  test-oauth-token\n", "", "", BoxStyle{})
	lines, err := box.GetStats(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if !called || len(lines) != 1 || !strings.Contains(lines[0], "Go") {
		t.Fatalf("unexpected stats: %v (called=%v)", lines, called)
	}
}
