package domain

type Book struct {
	Author         string `json:"author"`
	Title          string `json:"title"`
	PriceCents     int    `json:"price_cents"`
	Copies         int    `json:"copies"`
	IsAvailable    bool   `json:"is_available"`
	InventoryCents int    `json:"inventory_cents"`
}

func GetBooks(repoBooks []*Book) []*Book {
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
