package book

import (
	"github.com/Superm4n97/book-server-backend/utils"
	"net/http"
)

// Book defines the information of the book
type Book struct {
	Name    string           `json:"name"`
	ISBN    string           `json:"isbn"`
	Authors utils.AuthorList `json:"authors"`
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("list of all books"))

	//if err := database.AddAuthor(author.Author{
	//	Name:  "rasel",
	//	Email: "rasel@gmail.com",
	//}); err != nil {
	//	fmt.Println(fmt.Scanf("ERROR: %s", err.Error()))
	//}

	//err := database.GetAuthor("rasel")
	//if err != nil {
	//	fmt.Println("ERR: ", err.Error())
	//}

	w.WriteHeader(http.StatusOK)
}
