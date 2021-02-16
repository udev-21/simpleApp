package models

import "time"

type Author struct {
	ID        uint      `gorm:"not null" json:"id"`
	FirstName string    `gorm:"type:varchar(100)" json:"first_name"`
	LastName  string    `gorm:"type:varchar(100)" json:"last_name"`
	Email     string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Age       uint8     `gorm:"type:integer" json:"age"`
	BirthDate time.Time `gorm:"type:timestamp" json:"birth_date"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}
