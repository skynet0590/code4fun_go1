package handlers

import (
	"html/template"
	"github.com/go-chi/chi"
	"net/http"
)

type (
	Blog struct {
		Name	string
		Content	string
	}
)


func StartRouting(r *chi.Mux) (err error) {
	funcs := template.FuncMap {
		"convertToHTML": func(str string) template.HTML {
			return template.HTML(str)
		},
	}
	var userTpl *template.Template
	userTpl, err = template.New("user").Funcs(funcs).ParseFiles("./web/tmpl/layout.html")
	if err != nil {
		return
	}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		blogs := []Blog{
			{Name:"Golang is great", Content: "<h1>Golang</h1> Golang Golang Golang"},
			{Name:"Javascript is awesome", Content: "<h3>Javascript</h3><p>Is stronger day by day</p>"},
		}
		userTpl.ExecuteTemplate(w,"layout",  blogs)
	})

	r.Route("/sub", func(r chi.Router) {
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Sub router"))
		})
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from sub router"))
		})
	})
	return
}
