package shorten_url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		want    string
		counter int
	}{
		{
			name:    "Base case",
			counter: 1000,
			want:    "Op",
		},
		{
			name:    "Random case 1",
			counter: 263493931,
			want:    "RKjTp",
		},
		{
			name:    "Random case 2",
			counter: 1073741824,
			want:    "aaaaab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setCounterForTests(tt.counter)
			got := generateShortURL()
			assert.Equal(t, tt.want, got)
		})
	}
}
