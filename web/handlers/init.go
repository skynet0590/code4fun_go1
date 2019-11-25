package handlers

import (
	"github.com/skynet0590/code4fun_go1/web/helper"
	"html/template"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *helper.TmplHelper

type (
	Blog struct {
		Name	string
		Content	string
	}
	Map map[string]interface{}
)


func StartRouting(r *chi.Mux) (err error) {
	funcs := template.FuncMap {
		"html": func(str string) template.HTML {
			return template.HTML(str)
		},
	}
	tpl, err = helper.NewTPL(helper.TmplConfig{
		Name:        "webHTML",
		Dir:         "./web/tmpl/",
		Suffix:      "html",
		ProcessData: nil,
		FuncMap:     funcs,
	})
	if err != nil {
		return
	}
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "web/public")
	err = helper.FileServer(r, "/web/public", http.Dir(filesDir))
	if err != nil {
		return
	}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		blogs := []Blog{
			{Name:"Golang is great", Content: "<h1>Golang</h1> Golang Golang Golang"},
			{Name:"Javascript is awesome", Content: "<h3>Javascript</h3><p>Is stronger day by day</p>"},
		}
		tpl.Render(w,r,"layout", Map{"blogs": blogs})
	})

	r.Route("/sub", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Sub router"))
		})
		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from sub router"))
		})
	})
	r.Route("/auth", auth)
	return
}

