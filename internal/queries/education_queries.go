// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package queries

import (
	"github.com/ldatb/cv-api/internal/models"
	"github.com/ldatb/cv-api/internal/utils"
)

// Get all education data
func GetAllEducation() ([]models.Education, error) {
	var allEducation []models.Education
	if result := utils.DB.Database.Find(&allEducation); result.Error != nil {
		utils.Logger.Errorf("Something went wrong fetching data from the db. Error: %v", result.Error)
		return nil, result.Error
	}

	return allEducation, nil
}

// Get an education data from the ID
func GetEducationFromID(queryID uint) (*models.Education, error) {
	education := &models.Education{ID: queryID}
	if result := utils.DB.Database.First(education); result.Error != nil {
		utils.Logger.Errorf("Something went wrong fetching data from the db. Error: %v", result.Error)
		return nil, result.Error
	}
	return education, nil
}
