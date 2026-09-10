package domain

type Book struct {
	Author     string
	Title      string
	PriceCents int
	Copies     int
}

func (b Book) IsAvailable() bool {
	return b.Copies > 0
}

func (b Book) InventoryCents() int {
	return b.PriceCents * b.Copies
}
