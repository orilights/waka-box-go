package wakabox

import (
	"context"
	"testing"
)

func TestASCIIBarChart(t *testing.T) {
	for _, test := range []struct {
		percent float64
		size    int
		want    string
	}{
		{0, 20, "--------------------"},
		{35.8, 20, "#######-------------"},
		{16.3, 20, "###-----------------"},
		{100, 20, "####################"},
		{-1, 5, "-----"},
		{101, 5, "#####"},
		{50, 0, ""},
		{50, -1, ""},
	} {
		if got := GenerateBarChart(context.Background(), test.percent, test.size, "ASCII"); got != test.want {
			t.Errorf("percent=%v size=%d: got %q, want %q", test.percent, test.size, got, test.want)
		}
	}
}
