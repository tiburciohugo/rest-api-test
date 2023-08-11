package handler

import (
	"github.com/tiburciohugo/rest-api-test/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLiteDB()
}
