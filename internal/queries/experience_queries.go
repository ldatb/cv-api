// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package queries

import (
	"github.com/ldatb/cv-api/internal/models"
	"github.com/ldatb/cv-api/pkg/utils"
)

// Get all experience data
func GetAllExperience() ([]models.Experience, error) {
	var allExperience []models.Experience
	if result := utils.DB.Database.Find(&allExperience); result.Error != nil {
		utils.Logger.Errorf("Something went wrong fetching data from the db. Error: %v", result.Error)
		return nil, result.Error
	}

	return allExperience, nil
}

// Get an experience data from the ID
func GetExperienceFromID(queryID uint) (*models.Experience, error) {
	experience := &models.Experience{ID: queryID}
	if result := utils.DB.Database.First(experience); result.Error != nil {
		utils.Logger.Errorf("Something went wrong fetching data from the db. Error: %v", result.Error)
		return nil, result.Error
	}
	return experience, nil
}
