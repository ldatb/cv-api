// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package utils

import (
	"fmt"

	"github.com/ldatb/cv-api/internal/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

// Creates a connection with the database, makes the migrations and insert data
func InitializeDatabase(logger *log.Logger, config models.DatabaseConfigs) *gorm.DB {
	// Load configuration
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT)

	// Open connection to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gorm_logger.Default.LogMode(gorm_logger.Silent),
	})
	if err != nil {
		logger.Fatalf("Unable to connect to database. Error: %v", err)
	}

	// Makes migrations
	if err := makeMigrations(db); err != nil {
		logger.Fatalf("Unable to make database migrations. Error: %v", err)
	}

	// Insert data into database
	if err := insertData(db); err != nil {
		logger.Fatalf("Unable to insert data. Error: %v", err)
	}

	return db
}

// Makes initial migrations based on the models
func makeMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Experience{}, &models.Projects{}, &models.Courses{}, &models.Education{})
	return err
}

// Inserts data into database
// This is by no means the correct way to do this, but it works for the purpose of this project
func insertData(db *gorm.DB) error {
	// Get all data
	experienceData := GetExperienceData()
	projectsData := GetProjectsData()
	coursesData := GetCoursesData()
	educationData := GetEducationData()

	// Insert experience data
	if result := db.Create(experienceData); result.Error != nil {
		return result.Error
	}

	// Insert projects data
	if result := db.Create(projectsData); result.Error != nil {
		return result.Error
	}

	// Insert courses data
	if result := db.Create(coursesData); result.Error != nil {
		return result.Error
	}

	// Insert education data
	if result := db.Create(educationData); result.Error != nil {
		return result.Error
	}

	return nil
}
