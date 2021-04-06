package blog

import (
	"blog-y/pkg/common/config"
	"blog-y/pkg/common/log"
	"gorm.io/gorm"
)

// log
var l *log.Logger

func WithLogger(logger *log.Logger) {
	l = logger
}

// config
var c *config.Config

func WithConfig(config *config.Config) {
	c = config
}

// MySQL
var db *gorm.DB

func WithMySQL(gormDB *gorm.DB) {
	db = gormDB
	// todo delete
	db.AutoMigrate(&User{}, &Post{})
}
