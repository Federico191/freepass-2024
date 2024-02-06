package entity

import (
	"gorm.io/gorm"
	"time"
)

type CandidateInformation struct {
	CandidateId uint      `gorm:"primaryKey"`
	Candidate   Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Vision      *string   `gorm:"size:255"`
	Mission     *string   `gorm:"size:255"`
	Achievement *string   `gorm:"size:255"`
	Experience  *string   `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt
}
