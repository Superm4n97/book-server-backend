package book

import (
	"fmt"
	"github.com/Superm4n97/book-server-backend/database"
	"net/http"
)

func AllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("list of all books"))

	if err := database.Query(); err != nil {
		fmt.Println(fmt.Scanf("ERROR: %s", err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
