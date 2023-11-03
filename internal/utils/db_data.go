// Copyright (c) 2023 Lucas de Ataides <ldatb@icloud.com>
//
// SPDX-License-Identifier: MIT
package utils

import (
	"time"

	"github.com/ldatb/cv-api/internal/models"

	pq "github.com/lib/pq"
)

// Get the experience data
func GetExperienceData() []*models.Experience {
	return []*models.Experience{
		{
			ID:       1,
			JobTitle: "Software Development Intern",
			Company:  "Encora Brazil Division",
			Details: pq.StringArray{
				"Performing unitary and automated tests for a distributed cloud platform.",
				"Responsible for helping the automation of the deployment of these distributed cloud platforms.",
			},
			CurrentJob: false,
			StartDate:  time.Date(2022, time.January, 10, 12, 0, 0, 0, time.UTC),
			EndDate:    time.Date(2022, time.August, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:       2,
			JobTitle: "Software Development Assistant",
			Company:  "Encora Brazil Division",
			Details: pq.StringArray{
				"Cloud Platform Software Engineer working on a conteinerized OpenStack solution.",
				"Collaborating with DevOps, QA and other teams to improve the software delivery process",
				"Development and maintenance of Debian packages utilizing the dpkg-buildpackage tool",
				"Responsible for writing, maintaining, and enhancing configuration files for Infrastructure as Code projects (IaaC).",
				"Responsible for maintenance and development of improvements for automated tests, continuous integration pipeline and automation of deployments on bare metal servers using Jenkins.",
				"Teamwork using Agile methodology to deliver different product features in different sprints, working together with different teams from different nationalities and cultures.",
			},
			CurrentJob: true,
			StartDate:  time.Date(2022, time.August, 15, 12, 0, 0, 0, time.UTC),
		},
	}
}

// Get the projects data
func GetProjectsData() []*models.Projects {
	return []*models.Projects{
		{
			ID:    1,
			Title: "Personal website",
			Details: pq.StringArray{
				"Development of a personal website containing my CV, posts and contact information",
				"Static website made with Go - Hugo",
				"Deployed using IaaC on AWS Infrastructure",
				"CI/CD Pipelines to automatically update both the website and the API",
			},
			LinkToRepo: "https://github.com/ldatb/website",
			Timestamp:  time.Date(2023, time.October, 1, 12, 0, 0, 0, time.UTC),
		},
	}
}

// Get the courses data
func GetCoursesData() []*models.Courses {
	return []*models.Courses{
		{
			ID:              1,
			Title:           "Containers & Kubernetes Essentials",
			OfferedBy:       "IBM",
			CertificateLink: "https://www.credly.com/badges/20571be2-15d5-4571-9ee0-725dec1087c2",
			Duration:        time.Hour * 6,
			Timestamp:       time.Date(2021, time.May, 31, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:              2,
			Title:           "Linux LPI Essentials",
			OfferedBy:       "Alura",
			CertificateLink: "https://cursos.alura.com.br/degree/certificate/5294ab09-9e1a-4bd1-aa45-752792af9a43",
			Duration:        time.Hour * 30,
			Timestamp:       time.Date(2022, time.October, 3, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:              3,
			Title:           "DevOps Formation",
			OfferedBy:       "Alura",
			CertificateLink: "https://cursos.alura.com.br/degree/certificate/8874a862-084a-49bf-8a59-33b4e1bb5ced",
			Duration:        time.Hour * 58,
			Timestamp:       time.Date(2022, time.November, 7, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:              4,
			Title:           "Software Engineering",
			OfferedBy:       "Alura",
			CertificateLink: "https://cursos.alura.com.br/degree/certificate/1482b9c6-5a51-49ba-8da9-ca765a9ddf59",
			Duration:        time.Hour * 86,
			Timestamp:       time.Date(2022, time.November, 8, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:              5,
			Title:           "Infra as a Code",
			OfferedBy:       "Alura",
			CertificateLink: "https://cursos.alura.com.br/degree/certificate/6e5989fa-2053-4613-a0da-ba683db2e640",
			Duration:        time.Hour * 46,
			Timestamp:       time.Date(2023, time.January, 26, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:              6,
			Title:           "Mathematics for Machine Learning",
			OfferedBy:       "Imperial College London",
			CertificateLink: "https://www.coursera.org/account/accomplishments/specialization/4A3RWFUTUFPW",
			Duration:        time.Hour * 40,
			Timestamp:       time.Date(2023, time.March, 27, 12, 0, 0, 0, time.UTC),
		},
	}
}

// Get the education data
func GetEducationData() []*models.Education {
	return []*models.Education{
		{
			ID:              1,
			DegreeType:      "Technology Degree",
			DegreeName:      "Software Development",
			CollegeName:     "UNIFBV Wyden",
			CollegeLocation: "Brazil",
			StartDate:       time.Date(2022, time.April, 14, 12, 0, 0, 0, time.UTC),
			EndDate:         time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC),
		},
	}
}
