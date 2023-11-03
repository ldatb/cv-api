// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package server

import (
	"fmt"

	"github.com/ldatb/cv-api/configs"
	"github.com/ldatb/cv-api/internal/routes"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Initialize the API
func Start(port int, logger *log.Logger, database *gorm.DB) {
	// Get Fiber configs
	config := configs.FiberConfig()

	// Define a new Fiber app
	app := fiber.New(config)

	// Routes
	routes.PublicRoutes(app, database)
	routes.RouteNotFound(app)

	// Start API
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		logger.Fatalf("Oops... Looks like the server failed to start! Reason: %v", err)
	}
}
