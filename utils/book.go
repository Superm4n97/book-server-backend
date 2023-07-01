package utils

// Book defines the information of the book
type Book struct {
	ID      string     `json:"_id,omitempty"`
	Name    string     `json:"name"`
	Genre   []string   `json:"genre"`
	Authors AuthorList `json:"authors"`
}
