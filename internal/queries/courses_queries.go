// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package queries

import (
	"github.com/ldatb/cv-api/internal/models"
	"github.com/ldatb/cv-api/pkg/utils"
)

// Get all courses data
func GetAllCourses() ([]models.Courses, error) {
	var allCourses []models.Courses
	if result := utils.DB.Database.Find(&allCourses); result.Error != nil {
		utils.Logger.Errorf("Something went wrong fetching data from the db. Error: %v", result.Error)
		return nil, result.Error
	}

	return allCourses, nil
}

// Get an courses data from the ID
func GetCoursesFromID(queryID uint) (*models.Courses, error) {
	courses := &models.Courses{ID: queryID}
	if result := utils.DB.Database.First(courses); result.Error != nil {
		utils.Logger.Errorf("Something went wrong fetching data from the db. Error: %v", result.Error)
		return nil, result.Error
	}
	return courses, nil
}
