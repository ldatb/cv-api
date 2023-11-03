// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package configs

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

// Configurations for the Fiber server
func FiberConfig() fiber.Config {
	return fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		GETOnly:     true,
	}
}
