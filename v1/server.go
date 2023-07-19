package main

import (
	"log"
	"net/http"
	"v1/oss"
)

func main() {
	oss.OssLogger.Println("Server start success !")
	http.HandleFunc("/objects/", oss.Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
