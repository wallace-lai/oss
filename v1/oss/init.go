package oss

import (
	"log"
	"os"
)

const logPath string = "/var/log/oss/msg.log"

var OssLogger *log.Logger

func init() {
	fd, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("INIT : open file failed %s", err)
		return
	}

	OssLogger = log.New(fd, "[debug]", log.Ldate|log.Ltime)
}
