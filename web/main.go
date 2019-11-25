package main

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/skynet0590/code4fun_go1/web/handlers"
	"github.com/skynet0590/code4fun_go1/web/models"
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	err := handlers.StartRouting(r)
	if err != nil {
		showErr(mainErr{"starting router", err})
		return
	}
	err = models.StartDatabase()
	if err != nil {
		showErr(mainErr{"starting database", err})
		return
	}
	fmt.Println("Starting serve on port 8000")
	err = http.ListenAndServe(":8000", r)
	showErr(mainErr{"starting http server", err})
}

type (
	mainErr struct {
		name string
		originErr error
	}
)

func (m mainErr) Error() string {
	return fmt.Sprintf("Error was happend when %s: %v", m.name, m.originErr)
}

func showErr(err error) {
	fmt.Println(err.Error())
	fmt.Println("Start server was fail. The process was forced to exit!")
}