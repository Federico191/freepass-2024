package entity

type Post struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Description string  `gorm:"not null"`
	PicUrl      *string `gorm:"size:255"`
}
