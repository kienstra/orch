package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/kienstra/orch/internal/domain"
)

func TestGetBooks(t *testing.T) {
	tt := []struct {
		name     string
		expected []*domain.Book
		params   []*domain.Book
	}{
		{
			name:     "empty",
			expected: make([]*domain.Book, 0),
			params:   make([]*domain.Book, 0),
		},
		{
			name: "no copy",
			expected: []*domain.Book{
				{Copies: 0, PriceCents: 2100, InventoryCents: 0, IsAvailable: false},
			},
			params: []*domain.Book{{Copies: 0, PriceCents: 2100}},
		},
		{
			name: "1 book",
			expected: []*domain.Book{
				{Copies: 10, PriceCents: 2100, InventoryCents: 21000, IsAvailable: true},
			},
			params: []*domain.Book{{Copies: 10, PriceCents: 2100}},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, domain.GetBooks(tc.params))
	}
}
