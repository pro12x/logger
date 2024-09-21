package logger

import (
	"log"
	"os"
)

var (
	LoggerInfo  *log.Logger
	LoggerError *log.Logger
	logFile     *os.File
	Info        = "\033[34m" // Blue
	Error       = "\033[31m" // Red
	Warn        = "\033[33m" // Yellow
	Reset       = "\033[0m"
)

const maxFileSize = 24 * 1024

func InitLogger() {
	var err error
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		err := os.Mkdir("log", 0755)
		if err != nil {
			log.Println("Error creating log directory: " + err.Error())
			return
		}
	}
	logFile, err = os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening file: " + err.Error())
		return
	}
	LoggerInfo = log.New(logFile, Info+"INFO: "+Reset, log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	LoggerError = log.New(logFile, Error+"ERROR: "+Reset, log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	LoggerError = log.New(logFile, Warn+"WARN: "+Reset, log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
}
