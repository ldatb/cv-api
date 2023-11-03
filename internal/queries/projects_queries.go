// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package queries

import (
	"github.com/ldatb/cv-api/internal/models"
	"github.com/ldatb/cv-api/internal/utils"
)

// Get all projects data
func GetAllProjects() ([]models.Projects, error) {
	var allProjects []models.Projects
	if result := utils.DB.Database.Find(&allProjects); result.Error != nil {
		utils.Logger.Errorf("Something went wrong fetching data from the db. Error: %v", result.Error)
		return nil, result.Error
	}

	return allProjects, nil
}

// Get an projects data from the ID
func GetProjectsFromID(queryID uint) (*models.Projects, error) {
	projects := &models.Projects{ID: queryID}
	if result := utils.DB.Database.First(projects); result.Error != nil {
		utils.Logger.Errorf("Something went wrong fetching data from the db. Error: %v", result.Error)
		return nil, result.Error
	}
	return projects, nil
}
