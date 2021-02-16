package main

import (
	"time"
)

type Author struct {
	ID        uint      `gorm:"not null;primarykey" json:"id"`
	FirstName string    `gorm:"type:varchar(100)" json:"first_name" validate:"required,alphanumunicode"`
	LastName  string    `gorm:"type:varchar(100)" json:"last_name" validate:"required,alphanumunicode"`
	Email     string    `gorm:"type:varchar(100);unique_index" json:"email" validate:"required,email"`
	Age       uint8     `gorm:"type:integer" json:"age"`
	BirthDate time.Time `gorm:"type:timestamp" json:"birth_date"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

type Book struct {
	ID          uint      `gorm:"not null;primarykey" json:"id"`
	Title       string    `gorm:"type:varchar(100)" json:"title" validate:"required,alpha,min=2,max=100"`
	Description string    `gorm:"type:varchar(12288)" json:"description" validate:"required,alphanum,max=2048"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
	Limit       int       `gorm:"-" json:"limit"`
	Offset      int       `gorm:"-" json:"offset"`
	Author      Author    `json:"author"`
	AuthorID    int
}

type Response struct {
	Ok     bool        `json:"ok"`
	Errors []Error     `json:"errors"`
	Result interface{} `json:"result"`
}

type Error struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
