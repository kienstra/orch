package domain

import "github.com/kienstra/orch/internal/repository"

type Book struct {
	Author         string
	Title          string
	PriceCents     int
	Copies         int
	IsAvailable    bool
	InventoryCents int
}

func GetBooks(repoBooks []*repository.Book) []*Book {
	books := make([]*Book, len(repoBooks))
	for i, book := range repoBooks {
		books[i] = &Book{
			Author:         book.Author,
			Title:          book.Title,
			PriceCents:     book.PriceCents,
			Copies:         book.Copies,
			InventoryCents: book.PriceCents * book.Copies,
			IsAvailable:    book.Copies > 0,
		}
	}

	return books
}
