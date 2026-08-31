package server

import (
	"github.com/kienstra/orch/internal/domain"
	"github.com/kienstra/orch/internal/params"
	"github.com/kienstra/orch/internal/repository"

	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Repository repository.Repository
}

func (h *Server) Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetBooks).Methods(http.MethodGet)
}

func (h *Server) GetBooks(responseWriter http.ResponseWriter, req *http.Request) {
	reqParams, err := params.GetBooks(req)
	if err != nil {
		h.handleError(responseWriter, err)

		return
	}

	repoBooks, err := h.Repository.GetBooks(req.Context(), reqParams)
	if err != nil {
		h.handleError(responseWriter, err)

		return
	}

	domainBooks := domain.GetBooks(repoBooks)
	respond(responseWriter, domainBooks)
}

func (h *Server) handleError(responseWriter http.ResponseWriter, _ error) {
	http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
}

func respond(responseWriter http.ResponseWriter, response any) {
	b := []byte(response)
	_, err := responseWriter.Write(b)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
