// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package routes

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ldatb/cv-api/internal/models"
	"github.com/ldatb/cv-api/internal/queries"
	"github.com/ldatb/cv-api/internal/utils"
)

// Transform a Projects model into a map
func parseProjectsToMap(projects *models.Projects) map[string]interface{} {
	return map[string]interface{}{
		"id":         projects.ID,
		"title":      projects.Title,
		"details":    projects.Details,
		"repository": projects.LinkToRepo,
		"timestamp":  fmt.Sprint(projects.Timestamp.Year(), projects.Timestamp.Month(), projects.Timestamp.Day()),
	}
}

// ProjectsController method returns information about work projects
func ProjectsController(c *fiber.Ctx) error {
	// Return single projects data
	if c.Params("id") != "" {
		// Transform the 'id' param to an uint
		queryID, err := strconv.ParseUint(c.Params("id"), 10, 32)
		if err != nil {
			return utils.ReturnFiberError(c, fiber.StatusBadGateway, err)
		}

		// Fetch projects by ID
		projects, err := queries.GetProjectsFromID(uint(queryID))
		if err != nil {
			return utils.ReturnFiberError(c, fiber.StatusBadRequest, err)
		}
		projectsAsMap := parseProjectsToMap(projects)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  projectsAsMap,
		})
	}

	// Return all projects
	// Fetch all projects data
	allProjects, err := queries.GetAllProjects()
	if err != nil {
		return utils.ReturnFiberError(c, fiber.StatusBadRequest, err)
	}

	// Parse all projects into an array
	var allProjectsArray []map[string]interface{}
	for _, projects := range allProjects {
		projectsAsMap := parseProjectsToMap(&projects)
		allProjectsArray = append(allProjectsArray, projectsAsMap)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  allProjectsArray,
	})
}
