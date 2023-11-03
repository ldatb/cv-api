// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package models

import (
	"time"

	pq "github.com/lib/pq"
)

// Work experience
type Experience struct {
	ID         uint `gorm:"primaryKey"`
	JobTitle   string
	Company    string
	Details    pq.StringArray `gorm:"type:text[]"`
	CurrentJob bool
	StartDate  time.Time
	EndDate    time.Time
}

// Personal projects
type Projects struct {
	ID         uint `gorm:"primaryKey"`
	Title      string
	Details    pq.StringArray `gorm:"type:text[]"`
	LinkToRepo string
	Timestamp  time.Time
}

// Non-college courses
type Courses struct {
	ID              uint `gorm:"primaryKey"`
	Title           string
	OfferedBy       string
	CertificateLink string
	Duration        time.Duration
	Timestamp       time.Time
}

// College degrees
type Education struct {
	ID              uint `gorm:"primaryKey"`
	DegreeType      string
	DegreeName      string
	CollegeName     string
	CollegeLocation string
	StartDate       time.Time
	EndDate         time.Time
}
