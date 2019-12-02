package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/skynet0590/code4fun_go1/web/helper"
	"github.com/skynet0590/code4fun_go1/web/models"
	"net/http"
	"time"
)

func Start(r chi.Router) {
	r.Use(apiHeader)
	r.Route("/auth", auth)
	r.Route("/blog", blog)
}

const hmacSampleSecret  = "hmacSampleSecret"

func generateJWT(user models.User) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID": user.ID,
		"Email": user.Email,
		"exp": time.Now().Add( 24* time.Hour).Unix(),
	})
	tokenString, err = token.SignedString([]byte(hmacSampleSecret))
	return
}

func parseJWT(tokenString string) (user models.User, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.ID,_ = claims["ID"].(uint)
		user.Email,_ = claims["Email"].(string)
	} else {
		err = fmt.Errorf("Unknow error with JWT")
	}
	return
}

func neededLoginMDW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("Authorized")
		user, err := parseJWT(token)
		if err != nil {
			helper.JsonError(w, http.StatusUnauthorized, err, "Mã truy cập không hợp lệ hoặc đã hết hạn")
			return
		}
		ctx = context.WithValue(ctx, "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func apiHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// r.Header
		next.ServeHTTP(w, r)
	})
}