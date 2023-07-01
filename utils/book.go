package utils

import "github.com/Superm4n97/book-server-backend/utils/author"

// Book defines the information of the book
type Book struct {
	Name    string            `json:"name"`
	ISBN    string            `json:"isbn"`
	Authors author.AuthorList `json:"authors"`
}
