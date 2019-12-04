package helper

import (
	"encoding/json"
	"net/http"
	"github.com/go-playground/form"
)

var (
	ErrDataInvalid = "Dữ liệu không hợp lệ"
	ErrInternalServerError = "Lỗi nội bộ. Vui lòng liên hệ admin"
	ErrNotFound = "Dữ liệu không tồn tại"
)

type (
	Map map[string]interface{}
)

func Json(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	body,_ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

func JsonError(w http.ResponseWriter, code int, err error, msg string)  {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	Json(w,code, Map{
		"error": errStr,
		"msg": msg,
	})
}

func ParseJson(r *http.Request, obj interface{}) error {
	err := json.NewDecoder(r.Body).Decode(obj)
	return err
}

func ParseForm(r *http.Request, obj interface{}) error {
	r.ParseForm()
	decoder := form.NewDecoder()
	err := decoder.Decode(obj, r.Form)
	return err
}