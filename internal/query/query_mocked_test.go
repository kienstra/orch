package query_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kienstra/orch/internal/domain"
	"github.com/kienstra/orch/internal/repository"
	"github.com/kienstra/orch/internal/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestQueryMocked(t *testing.T) {
	tt := []struct {
		name       string
		req        string
		repoReturn []*domain.Book
		want       int
	}{
		{
			name: "empty",
			req:  "/",
			want: http.StatusBadRequest,
		},
		{
			name: "one book",
			req:  "/?page=100",
			want: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		repo := repository.NewMockRepository(gomock.NewController(t))
		if tc.repoReturn != nil {
			repo.EXPECT().GetBooks(
				gomock.Any(),
				gomock.Any()).Return(tc.repoReturn, nil)
		}
		request := httptest.NewRequest(http.MethodGet, tc.req, nil)
		recorder := httptest.NewRecorder()
		app := &server.Server{Repository: repo}

		app.GetBooks(recorder, request)
		assert.Equal(t, tc.want, recorder.Code)
	}
}
