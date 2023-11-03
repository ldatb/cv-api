// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package utils

import (
	"fmt"
	"time"

	"github.com/ldatb/cv-api/internal/models"
	"github.com/ldatb/cv-api/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

// Initialize global connection
var DB models.DBInstance

// Creates a connection with the database
func ConnectToDB(config models.DatabaseConfigs) {
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
		Logger.Fatalf("Unable to connect to database. Error: %v", err)
	}

	// Configures connections
	sqlDB, err := db.DB()
	if err != nil {
		Logger.Fatalf("Unable to get raw SQL connection. Error: %v", err)
	}
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Create global connection
	DB = models.DBInstance{
		Database: db,
	}
}

// Makes the migrations and insert data
func InitializeDatabase(config models.DatabaseConfigs) {
	// Makes migrations
	if err := makeMigrations(DB.Database); err != nil {
		Logger.Fatalf("Unable to make database migrations. Error: %v", err)
	}

	// Insert data into database
	if err := insertData(DB.Database); err != nil {
		Logger.Fatalf("Unable to insert data. Error: %v", err)
	}
}

// Makes initial migrations based on the models
func makeMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Experience{},
		&models.Projects{},
		&models.Courses{},
		&models.Education{},
	)
	return err
}

// Inserts data into database
// This is by no means the correct way to do this, but it works for the purpose of this project
func insertData(db *gorm.DB) error {
	// Get all data
	experienceData := utils.GetExperienceData()
	projectsData := utils.GetProjectsData()
	coursesData := utils.GetCoursesData()
	educationData := utils.GetEducationData()

	// Invalidate all data first
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Experience{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Projects{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Courses{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Education{})

	// Insert experience data
	if insert := db.Create(experienceData); insert.Error != nil {
		return insert.Error
	}

	// Insert projects data
	if insert := db.Create(projectsData); insert.Error != nil {
		return insert.Error
	}

	// Insert courses data
	if insert := db.Create(coursesData); insert.Error != nil {
		return insert.Error
	}

	// Insert education data
	if insert := db.Create(educationData); insert.Error != nil {
		return insert.Error
	}

	return nil
}
