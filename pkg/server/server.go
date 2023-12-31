// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package server

import (
	"fmt"

	"github.com/ldatb/cv-api/configs"
	"github.com/ldatb/cv-api/internal/routes"
	"github.com/ldatb/cv-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// Initialize the API
func Start(port string) {
	// Get Fiber configs
	config := configs.FiberConfig()

	// Define a new Fiber app
	app := fiber.New(config)

	// Routes
	routes.PublicRoutes(app)
	routes.RouteNotFound(app)

	// Start API
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		utils.Logger.Fatalf("Oops... Looks like the server failed to start! Reason: %v", err)
	}
}
