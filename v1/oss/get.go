package oss

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.EscapedPath(), "/")
	getName := splitPath[len(splitPath) - 1]
	OssLogger.Println("Get Name : ", getName)

	file, err := os.Open("/tmp/objects/" + getName)
	if err != nil {
		OssLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	defer file.Close()
	io.Copy(w, file)
}
