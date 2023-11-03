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

// Transform a Education model into a map
func parseEducationToMap(education *models.Education) map[string]interface{} {
	return map[string]interface{}{
		"id":           education.ID,
		"degree_type":  education.DegreeType,
		"degree_name":  education.DegreeName,
		"college_name": education.CollegeName,
		"location":     education.CollegeLocation,
		"start_date":   fmt.Sprint(education.StartDate.Year(), education.StartDate.Month(), education.StartDate.Day()),
		"end_date":     fmt.Sprint(education.EndDate.Year(), education.EndDate.Month(), education.EndDate.Day()),
	}
}

// EducationController method returns information about work education
func EducationController(c *fiber.Ctx) error {
	// Return single education data
	if c.Params("id") != "" {
		// Transform the 'id' param to an uint
		queryID, err := strconv.ParseUint(c.Params("id"), 10, 32)
		if err != nil {
			return utils.ReturnFiberError(c, fiber.StatusBadGateway, err)
		}

		// Fetch education by ID
		education, err := queries.GetEducationFromID(uint(queryID))
		if err != nil {
			return utils.ReturnFiberError(c, fiber.StatusBadRequest, err)
		}
		educationAsMap := parseEducationToMap(education)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  educationAsMap,
		})
	}

	// Return all education
	// Fetch all education data
	allEducation, err := queries.GetAllEducation()
	if err != nil {
		return utils.ReturnFiberError(c, fiber.StatusBadRequest, err)
	}

	// Parse all education into an array
	var allEducationArray []map[string]interface{}
	for _, education := range allEducation {
		educationAsMap := parseEducationToMap(&education)
		allEducationArray = append(allEducationArray, educationAsMap)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  allEducationArray,
	})
}
