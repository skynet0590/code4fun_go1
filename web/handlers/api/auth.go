package api

import (
	"github.com/go-chi/chi"
	"github.com/skynet0590/code4fun_go1/web/helper"
	"github.com/skynet0590/code4fun_go1/web/models"
	"net/http"
)

func auth(r chi.Router) {
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		helper.ParseJson(r, &user)
		if msg, err := user.GetForLogin();err == nil {
			token,err := generateJWT(user)
			if err != nil {
				helper.JsonError(w, http.StatusBadRequest,err,  msg)
				return
			}
			helper.Json(w, http.StatusOK, helper.Map {
				"msg": "Đăng nhập thành công",
				"data": user,
				"access_token": token,
			})
		}else{
			helper.JsonError(w, http.StatusBadRequest,err,  msg)
		}
	})
	r.Options("/login", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.Header)
		helper.Json(w, http.StatusOK, helper.Map{"msg":"Ok"})
	})
}
