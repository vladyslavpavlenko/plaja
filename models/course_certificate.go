package models

import "time"

// CourseCertificate is the course_certificate model.
type CourseCertificate struct {
	ID        uint
	UserID    uint `gorm:"not null"`
	User      User
	CourseID  uint `gorm:"not null"`
	Course    Course
	CreatedAt time.Time
	UpdatedAt time.Time
}
