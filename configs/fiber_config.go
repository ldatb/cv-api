// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package configs

import "github.com/gofiber/fiber/v2"

// Configurations for the Fiber server
func FiberConfig() fiber.Config {
	return fiber.Config{
		GETOnly: true,
	}
}
