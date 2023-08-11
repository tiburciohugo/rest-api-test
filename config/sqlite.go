package config

import (
	"os"

	"github.com/tiburciohugo/rest-api-test/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Database file does not exist, creating it: %s", dbPath)
		// create the database file
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error opening database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Todo{})
	if err != nil {
		logger.Errorf("Error automigrating database: %v", err)
		return nil, err
	}

	return db, nil
}
