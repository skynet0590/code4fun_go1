package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/skynet0590/code4fun_go1/web/models"
	"net/http"
)

const (
	authSessionKey = "authSessionKey"
)

func CommonMdw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx.Value("user")
		session, _ := store.Get(r, "auth")
		if user,ok := session.Values[authSessionKey]; ok {
			ctx = context.WithValue(ctx, "user", user)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func auth(r chi.Router)  {
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		tpl.Render(w,r,"auth_login", nil)
	})
	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "auth")
		session.Values = make(map[interface{}]interface{})
		err := session.Save(r, w)
		if err == nil {
			http.Redirect(w,r, "/", http.StatusMovedPermanently)
		}else{
			fmt.Println("Error when logout: ", err)
			http.Redirect(w,r, "/", http.StatusMovedPermanently)
		}
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		user := models.User{
			Email: r.FormValue("Email"),
			Password: r.FormValue("Password"),
		}
		if msg, err := user.GetForLogin();err == nil {
			session, _ := store.Get(r, "auth")
			session.Values[authSessionKey] = user
			err = session.Save(r, w)
			if err != nil {
				responseJson(w, http.StatusInternalServerError, Map{
					"error": err.Error(),
					"msg": "Đăng nhập thất bại",
				})
				return
			}
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
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}