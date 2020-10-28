package vlog

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	ErrorLog      *logrus.Logger
	AccessLog     *logrus.Logger
	errLogFile    = "./tmp/log/error.log"
	accessLogFile = "./tmp/log/access.log"
)

func init() {
	initErrorLog()
	initAccessLog()
}

func initErrorLog() {
	ErrorLog = logrus.New()
	ErrorLog.SetFormatter(&logrus.JSONFormatter{})
	f, err := os.OpenFile(errLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	ErrorLog.SetOutput(f)
}

func initAccessLog() {
	AccessLog = logrus.New()
	AccessLog.SetFormatter(&logrus.JSONFormatter{})
	f, err := os.OpenFile(accessLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	AccessLog.SetOutput(f)
}
