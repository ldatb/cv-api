// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package configs

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Define the possible configurations here.
// Don't use the internal/models folder as that is reserved for DB models.
type Configs struct {
	API_PORT int
}

// Load the app's configurations with Viper
func AppConfig(logger *log.Logger) Configs {
	// Set the config file name, kind and path
	//viper.SetConfigName("config")
	//viper.SetConfigType("yml")
	//viper.AddConfigPath(".")

	// Enable Viper to read Environment Variables
	viper.AutomaticEnv()

	// Set default variables
	viper.SetDefault("API_PORT", 3000)

	// Try to load the config
	//if err := viper.ReadInConfig(); err != nil {
	//	logger.Errorf("Error reading config file, %v", err)
	//	os.Exit(1)
	//}

	var config Configs
	if err := viper.Unmarshal(&config); err != nil {
		logger.Errorf("Unable to decode config into struct, %v", err)
		os.Exit(1)
	}

	return config
}
