// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package routes

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ldatb/cv-api/internal/models"
	"github.com/ldatb/cv-api/internal/queries"
	"github.com/ldatb/cv-api/internal/utils"
)

// Transform a Courses model into a map
func parseCoursesToMap(courses *models.Courses) map[string]interface{} {
	return map[string]interface{}{
		"id":               courses.ID,
		"title":            courses.Title,
		"offered_by":       courses.OfferedBy,
		"certificate_link": courses.CertificateLink,
		"course_duration":  fmt.Sprintf("%d hours", courses.Duration/time.Hour),
		"timestamp":        fmt.Sprint(courses.Timestamp.Year(), courses.Timestamp.Month(), courses.Timestamp.Day()),
	}
}

// CoursesController method returns information about work courses
func CoursesController(c *fiber.Ctx) error {
	// Return single courses data
	if c.Params("id") != "" {
		// Transform the 'id' param to an uint
		queryID, err := strconv.ParseUint(c.Params("id"), 10, 32)
		if err != nil {
			return utils.ReturnFiberError(c, fiber.StatusBadGateway, err)
		}

		// Fetch courses by ID
		courses, err := queries.GetCoursesFromID(uint(queryID))
		if err != nil {
			return utils.ReturnFiberError(c, fiber.StatusBadRequest, err)
		}
		coursesAsMap := parseCoursesToMap(courses)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  coursesAsMap,
		})
	}

	// Return all courses
	// Fetch all courses data
	allCourses, err := queries.GetAllCourses()
	if err != nil {
		return utils.ReturnFiberError(c, fiber.StatusBadRequest, err)
	}

	// Parse all courses into an array
	var allCoursesArray []map[string]interface{}
	for _, courses := range allCourses {
		coursesAsMap := parseCoursesToMap(&courses)
		allCoursesArray = append(allCoursesArray, coursesAsMap)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  allCoursesArray,
	})
}
