package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/skynet0590/code4fun_go1/web/handlers"
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	err := handlers.StartRouting(r)
	if err != nil {
		fmt.Printf("Error was happend when starting router: %v", err)
		return
	}
	fmt.Println("Starting serve on port 8000")
	http.ListenAndServe(":8000", r)

}