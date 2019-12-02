package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func blog(r chi.Router)  {
	r.Use(neededLoginMDW)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// For get list of blog
		fmt.Println("Get  list of blogs")
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// For get list of blog
		fmt.Println("Create a new blog")
	})
	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		// For get list of blog
		fmt.Println("Update a blog")
	})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		// For get list of blog
		fmt.Println("Delete a blog")
	})
}
