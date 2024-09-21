package main

import (
	"github.com/pro12x/logger"
)

func main() {
	logger.Process()

	logger.CatchLog("Hello, World!", "info")
	logger.CatchLog("This is an error message", "error")
	logger.CatchLog("This is a warning message", "warn")
}
