package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kienstra/orch/internal/query"
	"github.com/kienstra/orch/internal/repository"
)

type Server struct {
	Repository repository.Repository
}

func (s *Server) Serve(port string) error {
	router := mux.NewRouter()
	router.HandleFunc("/books", s.GetBooks).Methods(http.MethodGet)

	httpServer := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return httpServer.ListenAndServe()
}

// 1. Imperative shell ---------------------------v
func (s *Server) GetBooks(w http.ResponseWriter, req *http.Request) {
	// 2. Functional core
	input, err := query.GetBooks(req)
	if err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	// 3. Functional core: construct SQL query inside repository
	// 4. Imperative shell: query DB
	repoBooks, err := s.Repository.GetBooks(req.Context(), input)
	if err != nil {
		handleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// 5. Functional core
	res := ToBookResponses(repoBooks)

	// 6. Imperative shell.
	Respond(w, res)
}

func handleError(w http.ResponseWriter, message string, errorCode int) {
	http.Error(w, message, errorCode)
}

func Respond(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
