// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/ldatb/cv-api/internal/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group
	route := a.Group("/")

	// Define all routes
	route.Get("/", controllers.LandingPage)
}
