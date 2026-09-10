package server

import "github.com/kienstra/orch/internal/domain"

type bookResponse struct {
	Author         string `json:"author"`
	Title          string `json:"title"`
	PriceCents     int    `json:"price_cents"`
	Copies         int    `json:"copies"`
	IsAvailable    bool   `json:"is_available"`
	InventoryCents int    `json:"inventory_cents"`
}

func ToBookResponses(domainBooks []*domain.Book) []bookResponse {
	response := make([]bookResponse, len(domainBooks))
	for i, book := range domainBooks {
		response[i] = bookResponse{
			Author:         book.Author,
			Title:          book.Title,
			PriceCents:     book.PriceCents,
			Copies:         book.Copies,
			IsAvailable:    book.IsAvailable(),
			InventoryCents: book.InventoryCents(),
		}
	}

	return response
}
