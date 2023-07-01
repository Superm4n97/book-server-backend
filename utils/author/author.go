package author

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthorList struct {
	Authors []Author `json:"authors"`
}
