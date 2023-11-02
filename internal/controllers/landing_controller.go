// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package routes

import "github.com/gofiber/fiber/v2"

// LandingPage method returns the information about the API
func LandingPage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Welcome to Lucas' CV API!",
	})
}
