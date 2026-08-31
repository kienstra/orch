package main

import (
	"log"

	"github.com/kienstra/orch/internal/repository"
	"github.com/kienstra/orch/internal/server"
)

func main() {
	// Only need to inject repository. The domain is a pure function.
	repo, err := repository.NewRepositoryV1()
	if err != nil {
		log.Fatal(err)

		return
	}

	s := &server.Server{Repository: repo}
	s.Serve()
}
