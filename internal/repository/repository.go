package repository

import (
	"context"

	"github.com/kienstra/orch/internal/params"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Repository interface {
	GetBooks(context.Context, *params.Book) ([]*Book, error)
}

type RepositoryV1 struct {
	Db *sqlx.DB
}

type Book struct {
	Author     string
	Title      string
	Copies     int
	PriceCents int
}

func NewRepositoryV1() (*RepositoryV1, error) {
	db, err := sqlx.Open("sqlite", "orch.db")
	if err != nil {
		return nil, err
	}

	defer func() { _ = db.Close() }()
	return &RepositoryV1{db}, nil
}

func (r *RepositoryV1) GetBooks(ctx context.Context, params *params.Book) ([]*Book, error) {
	query, args := BooksSql(params)

	rows, err := r.Db.NamedQueryContext(ctx, query, args)
	defer func() { rows.Close() }()

	var books []*Book
	for rows.Next() {
		var book Book
		err = rows.StructScan(&book)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}

func BooksSql(params *params.Book) (string, map[string]any) {
	return `SELECT title, author
		FROM books
		WHERE author = :author
		OFFSET :offset
		LIMIT :limit`, map[string]any{
			"author": params.Author,
			"offset": (params.Page - 1) * params.Limit,
			"limit":  params.Limit,
		}
}
