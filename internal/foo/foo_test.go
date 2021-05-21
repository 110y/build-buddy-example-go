package foo_test

import (
	"testing"

	"github.com/110y/build-buddy-example-go/internal/foo"
)

func TestSum(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		x        int
		y        int
		expected int
	}{
		"1 + 1": {
			x:        1,
			y:        1,
			expected: 2,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if expected, actual := test.expected, foo.Sum(test.x, test.y); actual != expected {
				t.Errorf("expected: %d, actual: %d", expected, actual)
			}
		})
	}
}
