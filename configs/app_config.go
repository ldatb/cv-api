// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package configs

import (
	"log"

	"github.com/spf13/viper"
)

// Define the possible configurations here.
// Don't use the internal/models folder as that is reserved for DB models.
type Configs struct {
	API_PORT   int
	LOG_LEVEL  string
	LOG_OUTPUT string
}

// Defaults values for the configs
const (
	DEFAULT_API_PORT   = 3000
	DEFAULT_LOG_LEVEL  = "info"
	DEFAULT_LOG_OUTPUT = ""
)

// Load the app's configurations with Viper
func AppConfig() Configs {
	// Set the config file name, kind and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Enable Viper to read Environment Variables
	viper.AutomaticEnv()

	// Set default variables
	viper.SetDefault("API_PORT", DEFAULT_API_PORT)
	viper.SetDefault("LOG_LEVEL", DEFAULT_LOG_LEVEL)
	viper.SetDefault("LOG_OUTPUT", DEFAULT_LOG_OUTPUT)

	// Try to load the config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	// Parse configs into the struct
	var config Configs
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}

	return config
}
