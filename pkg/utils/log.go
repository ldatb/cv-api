// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package utils

import (
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Transforms a log level string into a logrus log level
func getLogLevel(level string, logger *log.Logger) log.Level {
	logLevels := map[string]log.Level{
		"trace": log.TraceLevel,
		"debug": log.DebugLevel,
		"info":  log.InfoLevel,
		"warn":  log.WarnLevel,
		"error": log.ErrorLevel,
		"fatal": log.FatalLevel,
		"panic": log.PanicLevel,
	}

	logLevel, ok := logLevels[strings.ToLower(level)]
	if !ok {
		// Defaults the log level to info, but send a warning message
		logger.Warnf("Unkown requested log level: %s, defaulting to info level", level)
		return logLevels["info"]
	}

	return logLevel
}

// If the log output is a file, create it
func createLogFile(logFile string, logger *log.Logger) io.Writer {
	// Create file
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("Unable to read / write to file %s", logFile)
	}
	defer file.Close()

	// Create writter of file
	return io.MultiWriter(os.Stdout, file)
}

// Create a global instance of the logger
var Logger *log.Logger

// Creates an instance of a logrus logger
func CreateLogger(logLevel string, logOutput string) {
	Logger = log.New()

	// Log level config
	loggerLogLevel := getLogLevel(logLevel, Logger)
	Logger.Infof("Setting log level to: %s", loggerLogLevel)
	Logger.SetLevel(loggerLogLevel)

	// Log output config
	if logOutput != "" {
		writter := createLogFile(logOutput, Logger)
		Logger.Infof("Setting log output to %s", logOutput)
		Logger.SetOutput(writter)
	} else {
		Logger.Debug("Setting log output to stdout")
		Logger.SetOutput(os.Stdout)
	}

	// Log formatter configs
	Logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
