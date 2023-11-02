// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Creates an instance of a logrus logger
func CreateLogger() *log.Logger {
	logger := log.New()

	// Logger configs
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	return logger
}
