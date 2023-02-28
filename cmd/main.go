package main

import "forum/pkg/logging"

func main() {
	logger := logging.GetLoggerInstance()
	logger.PrintInfo("logger initialized")
}
