# Logger Library
This project is a simple logging utility written in Go. It provides functionality to log messages with different severity levels (info, error, warn) and supports log file rotation.  

### Features
- Log messages with different severity levels: info, error, and warn.
- Automatically rotate log files when they reach a specified size.
- Initialize loggers and create log directories if they do not exist.
- Display a welcome message when a new log file is created.

### Installation
You can install the logger library using the following command:
```shell
go install github.com/pro12x/logger/cmd/logger@v1.0.0
```

### Usage
1. Initialize the logger and start logging messages:  
```go
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
```

2. Run your application:  
go run cmd/logger/main.go

### Functions
initLogger()
Initializes the loggers and creates the log directory if it does not exist. Calls the welcome function if a new log file is created.  
Process()
Initializes the logger, starts logging, and sets up signal handling for graceful shutdown.  
CatchLog(message, nature string)
Logs a message with the specified severity level (info, error, warn).  
rotateLogFile()
Rotates the log file when it reaches the specified size and initializes a new log file.  
cleanUp()
Closes the log file and performs cleanup operations.  
welcome()
Logs a welcome message when a new log file is created.  
Close()
Logs a closing message and performs final cleanup operations.  
Configuration
maxFileSize: The maximum size of the log file before it is rotated. Default is 24 KB.
_error, _warn, _reset: ANSI color codes for formatting log messages.

### License
This project is licensed under the MIT License. See the LICENSE file for details.