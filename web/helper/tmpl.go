package helper

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type (
	TmplHelper struct {
		tmpl   *template.Template
		config TmplConfig
	}
	TmplConfig struct {
		Name 	string
		Dir      string
		Suffix   string
		ProcessData func(*http.Request, map[string]interface{},) map[string]interface{}
		FuncMap template.FuncMap
		validFiles []string
	}
)

var (
	//cacheKey = ""
)

func NewTPL(config TmplConfig) (tmpl *TmplHelper, err error) {
	tmpl = &TmplHelper{}
	validFiles := []string{}
	err = filepath.Walk(config.Dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, "."+config.Suffix) {
			validFiles = append(validFiles, path)
		}
		return nil
	})
	if err != nil {
		return
	}
	/*config.FuncMap["uuid"] = func() string {
		return cacheKey
	}*/
	config.validFiles = validFiles
	tmpl.tmpl, err = template.New(config.Name).Funcs(config.FuncMap).ParseFiles(validFiles...)
	tmpl.config = config
	return
}

func (t TmplHelper) Render(wr io.Writer, r *http.Request, name string, data map[string]interface{}) {
	if t.config.ProcessData != nil {
		data = t.config.ProcessData(r,data)
	}
	err := t.tmpl.ExecuteTemplate(wr, name, data)
	if err != nil {
		fmt.Println(err)
	}
}
