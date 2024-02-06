package entity

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement"`
	Avatar              *string   `gorm:"size:255"`
	Username            string    `gorm:"size:20;not null"`
	User                User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BirthDate           time.Time `gorm:"type:date;not null"`
	IsVoted             *bool     `gorm:"default:false"`
	VotedElectionNumber *uint
	Candidate           []Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt           time.Time   `gorm:"autoCreateTime:milli"`
	UpdatedAt           time.Time   `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt           gorm.DeletedAt
}
