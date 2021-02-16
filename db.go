package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "udev21:udev21@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErr(err)
}
