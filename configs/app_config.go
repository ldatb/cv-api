// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package configs

import (
	"log"

	"github.com/ldatb/cv-api/internal/models"
	"github.com/spf13/viper"
)

// Defaults values for the configs
const (
	DEFAULT_API_PORT    = "3000"
	DEFAULT_LOG_LEVEL   = "info"
	DEFAULT_LOG_OUTPUT  = ""
	DEFAULT_DB_HOST     = "localhost"
	DEFAULT_DB_PORT     = "5432"
	DEFAULT_DB_USER     = "gorm"
	DEFAULT_DB_PASSWORD = "S3cretP@ss"
	DEFAULT_DB_NAME     = "api_db"
)

// Load the app's configurations with Viper
func AppConfig() models.Configs {
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
	viper.SetDefault("DB_HOST", DEFAULT_DB_HOST)
	viper.SetDefault("DB_PORT", DEFAULT_DB_PORT)
	viper.SetDefault("DB_USER", DEFAULT_DB_USER)
	viper.SetDefault("DB_PASSWORD", DEFAULT_DB_PASSWORD)
	viper.SetDefault("DB_NAME", DEFAULT_DB_NAME)

	// Try to load the config
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Error reading config file: %v. Using default configurationsdoc", err)
	}

	// Parse configs into the struct
	var config models.Configs
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}

	return config
}

// From all the configs, fetch only those related to the databse
func GetDatabaseConfigs(allConfigs models.Configs) models.DatabaseConfigs {
	return models.DatabaseConfigs{
		DB_HOST:     allConfigs.DB_HOST,
		DB_PORT:     allConfigs.DB_PORT,
		DB_USER:     allConfigs.DB_USER,
		DB_PASSWORD: allConfigs.DB_PASSWORD,
		DB_NAME:     allConfigs.DB_NAME,
	}
}
