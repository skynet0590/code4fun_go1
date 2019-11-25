package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/skynet0590/code4fun_go1/web/models"
	"net/http"
)

func auth(r chi.Router)  {
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		tpl.Render(w,r,"auth_login", nil)
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		user := models.User{
			Email: r.FormValue("Email"),
			Password: r.FormValue("Password"),
		}
		if msg, err := user.GetForLogin();err == nil {
			responseJson(w, http.StatusOK, Map{
				"msg": "Đăng nhập thành công",
				"data": user,
			})
		}else{
			responseJson(w, http.StatusBadRequest, Map{
				"error": err.Error(),
				"msg": msg,
			})
		}
	})
}

func responseJson(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	body,_ := json.Marshal(data)
	// w.Header().Add("")
	w.Write(body)
}