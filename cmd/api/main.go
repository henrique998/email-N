package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		w.Write([]byte(param))
	})

	r.Get("/{product}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "product")
		w.Write([]byte("O nome do produto é: " + param))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "success"}
		render.JSON(w, r, obj)
	})

	http.ListenAndServe(":3333", r)
}
