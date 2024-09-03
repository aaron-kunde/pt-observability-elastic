package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	log "pt.observability.elastic/app4/internal/logging"
)

type DataRepository struct {
	db *gorm.DB
}

var repository = DataRepository{
	db: initMySqlDBSession(),
}

func initMySqlDBSession() *gorm.DB {
	username := os.Getenv("GORM_DATASOURCE_USERNAME")

	if username == "" {
		username = "root"
	}
	password := os.Getenv("GORM_DATASOURCE_PASSWORD")

	if password == "" {
		password = "root"
	}
	url := os.Getenv("GORM_DATASOURCE_URL")

	if url == "" {
		url = "tcp(127.0.0.1:3306)/app4?charset=utf8mb4&parseTime=True&loc=Local"
	}
	dsn := fmt.Sprintf("%s:%s@%s", username, password, url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error(err)
	}
	return db
}

func Save(entity DataEntity) {
	result := repository.db.Create(&entity)

	if result.Error != nil {
		log.Error(result.Error)
	}
}
