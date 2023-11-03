// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package main

import (
	"github.com/ldatb/cv-api/configs"
	"github.com/ldatb/cv-api/internal/server"
	"github.com/ldatb/cv-api/internal/utils"
)

func main() {
	// Load configs
	config := configs.AppConfig()
	dbConfig := configs.GetDatabaseConfigs(config)

	// Create a logger
	logger := utils.CreateLogger(config.LOG_LEVEL, config.LOG_OUTPUT)

	// Initialize the database: Make migrations and insert data
	db := utils.InitializeDatabase(logger, dbConfig)

	// Start the API
	server.Start(config.API_PORT, logger, db)
}
