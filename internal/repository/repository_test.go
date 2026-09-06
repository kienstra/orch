package repository_test

import (
	"testing"

	params "github.com/kienstra/orch/internal/query"
	"github.com/kienstra/orch/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestBooksSql(t *testing.T) {
	tt := []struct {
		name         string
		params       *params.Book
		expectedArgs map[string]any
	}{
		{
			name:   "empty",
			params: &params.Book{},
			expectedArgs: map[string]any{
				"author": "",
				"limit":  0,
				"offset": 0,
			},
		},
		{
			name:   "page 1",
			params: &params.Book{Page: 1},
			expectedArgs: map[string]any{
				"author": "",
				"limit":  0,
				"offset": 0,
			},
		},
		{
			name:   "page 1",
			params: &params.Book{Page: 1, Limit: 10},
			expectedArgs: map[string]any{
				"author": "",
				"limit":  10,
				"offset": 0,
			},
		},
		{
			name:   "page and limit both over 1",
			params: &params.Book{Page: 2, Limit: 10},
			expectedArgs: map[string]any{
				"author": "",
				"limit":  10,
				"offset": 10,
			},
		},
	}

	for _, tc := range tt {
		_, actualArgs := repository.BooksSql(tc.params)
		assert.Equal(t, tc.expectedArgs, actualArgs)
	}
}
