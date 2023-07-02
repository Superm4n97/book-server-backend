package book

import (
	"encoding/json"
	"fmt"
	"github.com/Superm4n97/book-server-backend/database"
	"github.com/Superm4n97/book-server-backend/utils"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := chi.URLParam(r, "name")

	book, err := database.GetBook(name)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book utils.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = database.GetBook(book.Name)

	if err == mongo.ErrNoDocuments {
		err = database.AddBook(&book)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write([]byte(fmt.Sprintf("Book %s added successfully", book.Name)))
		w.WriteHeader(http.StatusOK)
		return
	}
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Book %s already exists", book.Name)))
	w.WriteHeader(http.StatusOK)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := chi.URLParam(r, "name")

	_, err := database.GetBook(name)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Book %s does not exists", name)))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = database.DeleteBook(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("Book %s successfully deleted", name)))
	return
}
