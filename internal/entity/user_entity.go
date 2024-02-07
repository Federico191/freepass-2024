package entity

import (
	"time"
)

type User struct {
	Username  string    `gorm:"primaryKey;size:20"`
	Email     string    `gorm:"size:50;not null;unique"`
	FullName  string    `gorm:"size:255;not null"`
	Password  string    `gorm:"size:60;not null"`
	IsAdmin   *bool     `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli"`
}
