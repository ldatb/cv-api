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
	config := configs.AppConfig()
	logger := utils.CreateLogger(config.LOG_LEVEL, config.LOG_OUTPUT)
	server.Start(config.API_PORT, logger)
}
