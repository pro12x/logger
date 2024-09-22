package main

import (
	"flag"
	"os"

	"github.com/pro12x/logger"
)

const version = "1.0.0"

func main() {
	versionFlag := flag.Bool("version", false, "Print the version of the logger")
	helpFlag := flag.Bool("help", false, "Print the help message for the logger")
	flag.Parse()

	if *helpFlag {
		logger.Help()
		return
	}

	if *versionFlag {
		os.Stdout.WriteString("Logger version " + version + "\n")
		return
	}

	logger.Process()

	logger.CatchLog("Hello, World!", "info")
	logger.CatchLog("This is an error message", "error")
	logger.CatchLog("This is a warning message", "warn")
}
