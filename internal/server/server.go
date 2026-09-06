package server

import (
	"encoding/json"

	"net/http"

	"github.com/kienstra/orch/internal/domain"
	"github.com/kienstra/orch/internal/params"
	"github.com/kienstra/orch/internal/repository"

	"github.com/gorilla/mux"
)

type Server struct {
	Repository repository.Repository
}

func (s *Server) Serve(port string) error {
	router := mux.NewRouter()
	router.HandleFunc("/books", s.GetBooks).Methods(http.MethodGet)
	return http.ListenAndServe(":"+port, router)
}

// 1. Imperative shell ---------------------------v
func (s *Server) GetBooks(w http.ResponseWriter, req *http.Request) {
	// 2. Functional core
	bookParams, err := params.GetBooks(req)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	// 3. Imperative shell
	repoBooks, err := s.Repository.GetBooks(req.Context(), bookParams)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// 4. Functional core
	domainBooks := domain.GetBooks(repoBooks)

	// 5. Imperative shell.
	respond(w, domainBooks)
}

func handleError(w http.ResponseWriter, message string, errorCode int) {
	http.Error(w, message, errorCode)
}

func respond(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
