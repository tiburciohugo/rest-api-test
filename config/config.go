package config

import "gorm.io/gorm"

var (
	db *gorm.DB
	logger *Logger
)

func Initialize() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}