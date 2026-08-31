package server_test

import (
	"testing"

	"net/http/httptest"

	"github.com/kienstra/orch/internal/params"
	"github.com/kienstra/orch/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetBook(t *testing.T) {
	tt := []struct {
		name     string
		params   string
		expected map[string]any
		wantErr  bool
	}{
		{
			name:    "empty",
			params:  "/",
			wantErr: true,
		},
		{
			name:   "page 1",
			params: "/?author=Milton&page=1",
			expected: map[string]any{
				"author": "Milton",
				"offset": 0,
				"limit":  1000,
			},
			wantErr: false,
		},
		{
			name:   "page 1",
			params: "/?page=4&limit=500&author=Milton",
			expected: map[string]any{
				"author": "Milton",
				"offset": 1500,
				"limit":  500,
			},
			wantErr: false,
		},
	}

	for _, tc := range tt {
		bookParams, err := params.GetBooks(httptest.NewRequest("GET", tc.params, nil))
		isErr := err != nil
		assert.Equal(t, tc.wantErr, isErr)
		if tc.wantErr {
			continue
		}

		sql, dbArgs := repository.BooksSql(bookParams)
		assert.Equal(t, tc.expected, dbArgs)
		assert.Contains(t, sql, "FROM books")
	}
}
