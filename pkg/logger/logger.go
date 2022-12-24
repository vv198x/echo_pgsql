package logger

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"time"
	"userSL/pkg/cfg"
)

const logDir = "./log"
const logExp = ".log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var LogFile *os.File

func Start() {
	os.MkdirAll("./log", os.ModePerm)

	logFilePath := filepath.Join(
		logDir,
		time.Now().Format("06.01.02")+logExp)

	LogFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Dont create log file")
	}
	if cfg.Get().Debug {
		log.SetOutput(LogFile)
	}
}

func SendSyslogMail(msg string) {
	syslog, err := syslog.New(syslog.LOG_MAIL, "userSL")
	if err != nil {
		log.Println("Load syslog ", err)
	} else {
		log.SetOutput(syslog)
		log.Println(msg)
		log.SetOutput(LogFile)
	}
}
