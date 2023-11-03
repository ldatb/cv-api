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

// Transform an Experience model into a map
func parseExperienceToMap(experience *models.Experience) map[string]interface{} {
	experienceMap := map[string]interface{}{
		"id":             experience.ID,
		"job_title":      experience.JobTitle,
		"company":        experience.Company,
		"details":        experience.Details,
		"is_current_job": experience.CurrentJob,
		"start_date":     fmt.Sprint(experience.StartDate.Year(), experience.StartDate.Month(), experience.StartDate.Day()),
	}

	// Add end date if the job is not the current one
	if !experience.CurrentJob {
		experienceMap["end_data"] = fmt.Sprint(experience.EndDate.Year(), experience.EndDate.Month(), experience.EndDate.Day())
	}

	return experienceMap
}

// ExperienceController method returns information about work experience
func ExperienceController(c *fiber.Ctx) error {
	// Return single data experience
	if c.Params("id") != "" {
		// Transform the 'id' param to an uint
		queryID, err := strconv.ParseUint(c.Params("id"), 10, 32)
		if err != nil {
			return utils.ReturnFiberError(c, fiber.StatusBadGateway, err)
		}

		// Fetch experience by ID
		experience, err := queries.GetExperienceFromID(uint(queryID))
		if err != nil {
			return utils.ReturnFiberError(c, fiber.StatusBadRequest, err)
		}
		experienceAsMap := parseExperienceToMap(experience)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  experienceAsMap,
		})
	}

	// Return all experiences
	// Fetch all experience data
	allExperience, err := queries.GetAllExperience()
	if err != nil {
		return utils.ReturnFiberError(c, fiber.StatusBadRequest, err)
	}

	// Parse all experience into an array
	var allExperienceArray []map[string]interface{}
	for _, experience := range allExperience {
		experienceAsMap := parseExperienceToMap(&experience)
		allExperienceArray = append(allExperienceArray, experienceAsMap)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  allExperienceArray,
	})
}
