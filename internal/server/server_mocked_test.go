package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kienstra/orch/internal/domain"
	"github.com/kienstra/orch/internal/repository"
	"github.com/kienstra/orch/internal/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServerMocked(t *testing.T) {
	zero := float64(0)
	tt := []struct {
		name       string
		repoReturn []*domain.Book
		want       []map[string]any
	}{
		{
			name:       "empty",
			repoReturn: nil,
			want:       []map[string]any{},
		},
		{
			name: "one book",
			repoReturn: []*domain.Book{{
				Author: "Milton",
				Title:  "Paradise Lost",
			}},
			want: []map[string]any{{
				"author":          "Milton",
				"title":           "Paradise Lost",
				"copies":          zero,
				"price_cents":     zero,
				"inventory_cents": zero,
				"is_available":    false,
			}},
		},
	}

	for _, tc := range tt {
		repo := repository.NewMockRepository(gomock.NewController(t))
		repo.EXPECT().GetBooks(
			gomock.Any(),
			gomock.Any()).Return(tc.repoReturn, nil)
		request := httptest.NewRequest(http.MethodGet, "/books?author=Milton", nil)
		recorder := httptest.NewRecorder()
		app := &server.Server{Repository: repo}

		app.GetBooks(recorder, request)
		assert.Equal(t, http.StatusOK, recorder.Code)
		var got []map[string]any
		_ = json.Unmarshal(recorder.Body.Bytes(), &got)
		assert.Equalf(t, tc.want, got, "something")
	}
}
