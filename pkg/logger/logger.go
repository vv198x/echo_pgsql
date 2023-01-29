package logger

import (
	"log"
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
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Println("Log dir exist")
	}

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
