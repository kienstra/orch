package query_test

import (
	"net/http/httptest"
	"testing"

	params "github.com/kienstra/orch/internal/query"

	"github.com/stretchr/testify/assert"
)

func TestGetBook(t *testing.T) {
	tt := []struct {
		name     string
		req      string
		expected *params.Book
		wantErr  bool
	}{
		{
			name:     "empty",
			req:      "/",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "only page",
			req:      "/?page=100",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "only limit",
			req:      "/?limit=5",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "limit too low",
			req:      "/?author=Hemingway&limit=-1&page=100",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "limit too high",
			req:      "/?author=Hemingway&limit=1001&page=100",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "page too low",
			req:      "/?author=Hemingway&limit=1001&page=-1",
			expected: nil,
			wantErr:  true,
		},
		{
			name: "valid page and limit",
			req:  "/?author=Hemingway&limit=5&page=100",
			expected: &params.Book{
				Author: "Hemingway",
				Limit:  5,
				Page:   100,
			},
			wantErr: false,
		},
	}

	for _, tc := range tt {
		actual, err := params.GetBooks(httptest.NewRequest("GET", tc.req, nil))
		assert.Equal(t, tc.expected, actual, tc.name)
		isErr := err != nil

		if isErr && !tc.wantErr {
			t.Fatalf("unexpectedly got error %s", err.Error())
		}

		if !isErr && tc.wantErr {
			t.Fatalf("expected error, but got %v", err)
		}
	}
}
