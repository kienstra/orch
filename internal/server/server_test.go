package server_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/kienstra/orch/internal/domain"
	"github.com/kienstra/orch/internal/query"
	"github.com/kienstra/orch/internal/repository"
	"github.com/kienstra/orch/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestParamsToSql(t *testing.T) {
	tt := []struct {
		name     string
		query    string
		expected map[string]any
		wantErr  bool
	}{
		{
			name:    "empty",
			query:   "/",
			wantErr: true,
		},
		{
			name:  "page 1",
			query: "/?author=Milton&page=1",
			expected: map[string]any{
				"author": "Milton",
				"offset": 0,
				"limit":  1000,
			},
			wantErr: false,
		},
		{
			name:  "page 1",
			query: "/?page=4&limit=500&author=Milton",
			expected: map[string]any{
				"author": "Milton",
				"offset": 1500,
				"limit":  500,
			},
			wantErr: false,
		},
	}

	for _, tc := range tt {
		bookParams, err := query.GetBooks(httptest.NewRequest("GET", tc.query, nil))
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

func TestDomainToResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	server.Respond(recorder, server.ToBookResponses([]*domain.Book{
		{
			Author: "Milton",
			Title:  "Paradise Lost",
		},
	}))

	var got []map[string]any
	_ = json.Unmarshal(recorder.Body.Bytes(), &got)
	assert.Equal(t, "Milton", got[0]["author"])
}
