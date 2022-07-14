package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type m  struct{
	Error string `json:"error"`
	Method string `json:"method"`
	Status int `json:"status"`
}

func WrapErrorWithStatus( w http.ResponseWriter, err error, status string, httpStatus int)  {
	m := m{ Error: err.Error(), Method: status, Status: httpStatus}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset = utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff") //даем понять что ответ приходит в формате json
	w.WriteHeader(httpStatus)
	fmt.Fprintln(w, string(res))
}

func WrapOK(w http.ResponseWriter)  {
	var m = map[string]string{
		"status": "result",
		"message" : "created",
	}
	res, _:= json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w , string(res))
}
