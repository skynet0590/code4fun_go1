package helper

import (
	"encoding/json"
	"net/http"
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