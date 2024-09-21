package logger

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	logger      *log.Logger
	loggerInfo  *log.Logger
	loggerError *log.Logger
	loggerWarn  *log.Logger
	logFile     *os.File
	_info       = "\033[34m" // Blue
	_error      = "\033[31m" // Red
	_warn       = "\033[33m" // Yellow
	_reset      = "\033[0m"
)

const (
	fileName = "log/app.log"
)

const maxFileSize = 24 * 1024

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func initLogger() {
	var (
		err   error
		check bool
	)
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		err := os.Mkdir("log", 0755)
		if err != nil {
			log.Println("_error creating log directory: " + err.Error())
			return
		}
	}

	if fileExists(fileName) {
		check = true
	}

	logFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("_error opening file: " + err.Error())
		return
	}

	logger = log.New(logFile, "", 0)
	loggerInfo = log.New(logFile, _info+"INFO: "+_reset, log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	loggerError = log.New(logFile, _error+"ERROR: "+_reset, log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	loggerWarn = log.New(logFile, _warn+"WARN: "+_reset, log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	if !check {
		welcome()
	}
}

func Process() {
	initLogger()
	loggerInfo.Println("Logger started...")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		cleanUp()
		os.Exit(0)
	}()

	rotateLogFile()
}

func CatchLog(message, nature string) {
	switch nature {
	case "info":
		loggerInfo.Println(message)
	case "error":
		loggerError.Println(_error + message + _error)
	case "warn":
		loggerWarn.Println(_warn + message + _reset)
	case "none":
		logger.Println(_warn + message + _reset)
	default:
		loggerError.Println(_error + "Unknown log nature: " + nature + _reset)
	}
}

func rotateLogFile() {
	if stat, err := logFile.Stat(); err == nil {
		if stat.Size() >= maxFileSize {
			err := logFile.Close()
			if err != nil {
				log.Println("_error closing file: " + err.Error())
				return
			}

			newName := "log/app-" + time.Now().Format("2006-01-02_15-04-05") + ".log"
			err = os.Rename(fileName, newName)
			if err != nil {
				log.Println("_error renaming file: " + err.Error())
				CatchLog(_error+"_error renaming file: "+err.Error(), "error")
				return
			}

			CatchLog("CatchLog file rotated and file renamed to: "+newName, "info")

			initLogger()
		}
	}
}

func cleanUp() {
	if logFile != nil {
		Close()
		err := logFile.Close()
		if err != nil {
			CatchLog("_error closing file: "+err.Error(), "error")
			return
		}
	}
}

func welcome() {
	CatchLog(_warn+"██╗      ██████╗  ██████╗  ██████╗ ███████╗██████╗"+_reset, "none")
	CatchLog(_warn+"██║     ██╔═══██╗██╔════╝ ██╔════╝ ██╔════╝██╔══██╗"+_reset, "none")
	CatchLog(_warn+"██║     ██║   ██║██║  ███╗██║  ███╗█████╗  ██████╔╝"+_reset, "none")
	CatchLog(_warn+"██║     ██║   ██║██║   ██║██║   ██║██╔══╝  ██╔══██╗"+_reset, "none")
	CatchLog(_warn+"███████╗╚██████╔╝╚██████╔╝╚██████╔╝███████╗██║  ██║"+_reset, "none")
	CatchLog(_warn+"╚══════╝ ╚═════╝  ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝"+_reset, "none")
}

func Close() {
	CatchLog("Logger closed", "info")
	CatchLog("See you later!", "info")
	log.Printf("\n\n\n")
}
