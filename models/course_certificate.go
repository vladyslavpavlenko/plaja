package models

import "time"

// CourseCertificate is the course_certificate model.
type CourseCertificate struct {
	ID        uint
	AccountID uint `gorm:"not null"`
	Account   Account
	CourseID  uint `gorm:"not null"`
	Course    Course
	CreatedAt time.Time
	UpdatedAt time.Time
}