package models

import "time"

// Enrollment is the enrollment model.
type Enrollment struct {
	UserID         uint `gorm:"primaryKey;autoIncrement:false;not null"`
	User           User
	CourseID       uint `gorm:"primaryKey;autoIncrement:false;not null"`
	Course         Course
	Progress       uint
	StatusID       uint `gorm:"not null"`
	Status         EnrollmentStatus
	LastExerciseID uint `gorm:"not null"`
	LastExercise   CourseExercise
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
