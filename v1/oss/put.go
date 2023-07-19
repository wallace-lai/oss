package oss

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.EscapedPath(), "/")
	putName := splitPath[len(splitPath)-1]
	OssLogger.Println("Put Name : ", putName)

	file, err := os.Create("/tmp/objects/" + putName)
	if err != nil {
		OssLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()
	io.Copy(file, r.Body)
}
