package oss

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == http.MethodPut {
		OssLogger.Println("PUT method received !")
		put(w, r)
		return
	} else if method == http.MethodGet {
		OssLogger.Println("GET method received !")
		get(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
