package database

import (
	"github.com/jinzhu/gorm" // jinzhu - creator
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DbConn *gorm.DB
)