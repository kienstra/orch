package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kienstra/orch/internal/domain"
	"github.com/kienstra/orch/internal/query"
	_ "github.com/lib/pq"
)

type PgRepository struct {
	Db *sqlx.DB
}

type bookRow struct {
	Author     string `db:"author"`
	Title      string `db:"title"`
	Copies     int    `db:"copies"`
	PriceCents int    `db:"price_cents"`
}

type DbConfig struct {
	DbUser string
	DbPass string
	DbHost string
	DbName string
}

type Repository interface {
	GetBooks(context.Context, *query.Book) ([]*domain.Book, error)
}

func NewPgRepository(url string) (*PgRepository, error) {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PgRepository{Db: db}, nil
}

func (r *PgRepository) Close() {
	_ = r.Db.Close()
}

func (r *PgRepository) GetBooks(ctx context.Context, params *query.Book) ([]*domain.Book, error) {
	sql, args := BooksSql(params)

	rows, err := r.Db.NamedQueryContext(ctx, sql, args)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var books []*domain.Book
	for rows.Next() {
		var book bookRow
		err = rows.StructScan(&book)
		if err != nil {
			return nil, err
		}

		books = append(books, &domain.Book{
			Author:     book.Author,
			Title:      book.Title,
			Copies:     book.Copies,
			PriceCents: book.PriceCents,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func BooksSql(params *query.Book) (string, map[string]any) {
	return `SELECT title, author, price_cents, copies
		FROM books
		WHERE author = :author
		ORDER BY author
		OFFSET :offset
		LIMIT :limit`, map[string]any{
			"author": params.Author,
			"offset": (params.Page - 1) * params.Limit,
			"limit":  params.Limit,
		}
}
