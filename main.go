package main

import (
	"log"
	"os"

	"github.com/kienstra/orch/internal/repository"
	"github.com/kienstra/orch/internal/server"
)

func main() {
	repo, err := repository.NewPgRepository(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer repo.Close()

	s := &server.Server{Repository: repo}
	err = s.Serve(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
