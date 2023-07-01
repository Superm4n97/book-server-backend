package utils

type Author struct {
	ID    string `json:"_id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthorList struct {
	Authors []Author `json:"authors"`
}
