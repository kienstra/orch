package query

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	Author string
	Limit  int
	Page   int
}

const (
	defaultLimit = 1000
	defaultPage  = 1
	minPage      = 1
	minLimit     = 1
	maxLimit     = 1000
)

func GetBooks(req *http.Request) (*Book, error) {
	pageQuery := req.URL.Query().Get("page")
	limitQuery := req.URL.Query().Get("limit")
	authorQuery := req.URL.Query().Get("author")

	if authorQuery == "" {
		return nil, errors.New("author is empty")
	}

	page, err := getPage(pageQuery)
	if err != nil {
		return nil, fmt.Errorf("invalid page: %w", err)
	}

	limit, err := getLimit(limitQuery)
	if err != nil {
		return nil, fmt.Errorf("invalid limit: %w", err)
	}

	return &Book{
		Author: authorQuery,
		Limit:  limit,
		Page:   page,
	}, nil
}

func getPage(pageQuery string) (int, error) {
	if pageQuery == "" {
		return defaultPage, nil
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		return 0, err
	}

	if page < minPage {
		return 0, fmt.Errorf("page of %d is less than the minimum of %d", page, minPage)
	}

	return page, nil
}

func getLimit(limitQuery string) (int, error) {
	if limitQuery == "" {
		return defaultLimit, nil
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		return 0, err
	}

	if limit < minLimit {
		return 0, fmt.Errorf("limit of %d is below minimum of %d", limit, minLimit)
	}

	if limit > maxLimit {
		return 0, fmt.Errorf("limit of %d is above maximum of %d", limit, maxLimit)
	}

	return limit, nil
}
