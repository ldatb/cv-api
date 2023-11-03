// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package utils

import "github.com/gofiber/fiber/v2"

// Returns an error to Fiber
func ReturnFiberError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
