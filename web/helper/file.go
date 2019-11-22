package helper

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

func FileServer(r chi.Router, path string, root http.FileSystem) error {
	if strings.ContainsAny(path, "{}*") {
		return fmt.Errorf("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
	return nil
}