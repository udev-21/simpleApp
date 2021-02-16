package main

import "gorm.io/gorm"

var migrateFlag bool // run with migrate
var dbSeedFlag bool  // run with fill fake data
var db *gorm.DB      //global connection to database
