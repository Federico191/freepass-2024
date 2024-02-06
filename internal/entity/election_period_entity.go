package entity

import "time"

type ElectionPeriod struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	StartTime     time.Time `gorm:"not null"`
	EndTime       time.Time `gorm:"not null"`
	AdminUsername *string   `gorm:"size:20"`
	Admin         User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
