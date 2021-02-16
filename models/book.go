package models

import "time"

type Book struct {
	ID          uint      `gorm:"not null" json:"id"`
	Title       string    `gorm:"type:varchar(100)" json:"title"`
	Description string    `gorm:"type:varchar(12288)" json:"description"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
}
