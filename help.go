package logger

import "fmt"

func Help() {
	// Help is a function that prints out the help message for the logger package.
	//
	// Usage:
	// 	logger [flags]
	//
	// Flags:
	//   -version
	//         Print the version of the logger
	//
	// Examples:
	// 	logger -version
	//
	// For more information, see
	fmt.Println(`Help is a function that prints out the help message for the logger package.

	Usage:
		logger [flags]

	Flags:
		-version: Print the version of the logger

		-help: Print the help message for the logger

		Examples:
			logger -version
			logger -help`)
}
