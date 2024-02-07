package entity

import (
	"gorm.io/gorm"
	"time"
)

type Candidate struct {
	ID             uint    `gorm:"primaryKey;autoIncrement"`
	ElectionNumber *uint   `gorm:"unique;index"`
	AccountId      uint    `gorm:"not null"`
	Account        Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LeaderId       *uint
	Candidate      *Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      time.Time  `gorm:"autoCreateTime:milli"`
	UpdatedAt      time.Time  `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt      gorm.DeletedAt
}
