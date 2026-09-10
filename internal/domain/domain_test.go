package domain_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kienstra/orch/internal/domain"
)

func TestIsAvailable(t *testing.T) {
	tt := []struct {
		name     string
		expected bool
		arg      int
	}{
		{
			name:     "negative",
			expected: false,
			arg:      -1,
		},
		{
			name:     "zero",
			expected: false,
			arg:      0,
		},
		{
			name:     "positive",
			expected: true,
			arg:      1,
		},
		{
			name:     "huge",
			expected: true,
			arg:      math.MaxInt,
		},
	}

	for _, tc := range tt {
		d := &domain.Book{Copies: tc.arg}
		assert.Equal(t, tc.expected, d.IsAvailable())
	}
}
