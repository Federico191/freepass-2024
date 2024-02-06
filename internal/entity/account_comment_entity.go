package entity

import (
	"gorm.io/gorm"
	"time"
)

type AccountComment struct {
	PostId      uint      `gorm:"primaryKey"`
	Post        Post      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CandidateId uint      `gorm:"primaryKey"`
	Candidate   Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccountId   uint      `gorm:"primaryKey"`
	Account     Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt
}
