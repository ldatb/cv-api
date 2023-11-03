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
	utils.CreateLogger(config.LOG_LEVEL, config.LOG_OUTPUT)

	// Connect to DB
	utils.ConnectToDB(dbConfig)

	// Initialize the database: Make migrations and insert data
	utils.InitializeDatabase(dbConfig)

	// Start the API
	server.Start(config.API_PORT)
}
