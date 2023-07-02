package router

import (
	"github.com/Superm4n97/book-server-backend/router/apis/v1/author"
	"github.com/Superm4n97/book-server-backend/router/apis/v1/book"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Pong"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func Register() *chi.Mux {
	rt := chi.NewRouter()
	rt.Get("/ping", pong)
	rt.Route("/apis/v1", func(r chi.Router) {
		r.Route("/book", func(r chi.Router) {
			r.Get("/{name}", book.AllBooks)
		})
		r.Route("/author", func(r chi.Router) {
			r.Get("/{name}", author.Get)
			r.Post("/", author.Add)
			r.Delete("/{name}", author.Delete)
		})
	})
	return rt
}
