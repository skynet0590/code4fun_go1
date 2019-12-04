package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/skynet0590/code4fun_go1/web/helper"
	"github.com/skynet0590/code4fun_go1/web/models"
	"net/http"
)

func blog(r chi.Router)  {
	r.Use(neededLoginMDW)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		filter := models.BlogFilter{
		}
		err := helper.ParseForm(r, &filter)
		if  err != nil {
			helper.JsonError(w, http.StatusBadRequest, err, helper.ErrDataInvalid)
			return
		}
		blogs, err := filter.GetList()
		if  err != nil {
			helper.JsonError(w, http.StatusInternalServerError, err, helper.ErrInternalServerError)
			return
		}
		helper.Json(w, http.StatusOK, blogs)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var blog models.Blog
		if err := helper.ParseJson(r, &blog); err != nil {
			helper.JsonError(w, http.StatusBadRequest, err, helper.ErrDataInvalid)
			return
		}
		if err := blog.Create(); err != nil {
			helper.JsonError(w, http.StatusInternalServerError, err, helper.ErrInternalServerError)
			return
		}
		helper.Json(w, http.StatusOK, blog)
	})
	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
		var blog, oldBlog models.Blog
		if err := helper.ParseJson(r, &blog); err != nil {
			helper.JsonError(w, http.StatusBadRequest, err, helper.ErrDataInvalid)
			return
		}
		oldBlog.ID = blog.ID
		if err := oldBlog.GetById(); err != nil {
			helper.JsonError(w, http.StatusNotFound, err, helper.ErrNotFound)
			return
		}
		oldBlog.Title = blog.Title
		oldBlog.Content = blog.Content
		if err := oldBlog.Update(); err != nil {
			helper.JsonError(w, http.StatusInternalServerError, err, helper.ErrInternalServerError)
			return
		}
		helper.Json(w, http.StatusOK, blog)
	})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		// For get list of blog
		fmt.Println("Delete a blog")
	})
}
