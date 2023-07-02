package author

import (
	"encoding/json"
	"fmt"
	"github.com/Superm4n97/book-server-backend/database"
	"github.com/Superm4n97/book-server-backend/utils"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/klog/v2"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authorName := chi.URLParam(r, "name")
	athr, err := database.GetAuthor(authorName)
	if err != nil {
		klog.Errorf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(athr)
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
	var athr utils.Author
	err := json.NewDecoder(r.Body).Decode(&athr)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = database.GetAuthor(athr.Name)
	if err == mongo.ErrNoDocuments {
		if err = database.AddAuthor(&athr); err != nil {
			_, _ = w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Author %s successfully added", athr.Name)))

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("Author %s already exists", athr.Name)))
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := chi.URLParam(r, "name")

	_, err := database.GetAuthor(name)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Author %s does not exists", name)))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = database.DeleteAuthor(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("Author %s successfully deleted", name)))
	return
}
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := database.ListAuthor()
	if err != nil {
		klog.Infof(err.Error())
	}
	return
}
