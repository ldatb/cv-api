// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package models

// All possible configs for the app
type Configs struct {
	API_PORT    string
	LOG_LEVEL   string
	LOG_OUTPUT  string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

// Only database configs
type DatabaseConfigs struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}
