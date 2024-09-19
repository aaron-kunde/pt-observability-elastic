package db

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	log "pt.observability.elastic/app4/internal/logging"
)

type DataRepository struct {
	db *gorm.DB
}

var (
	repository = DataRepository{
		db: initDBSession(),
	}
	tracer = otel.Tracer("")
)

func initDBSession() *gorm.DB {
	if os.Getenv("GORM_DRIVER") == "postgres" {
		return initPostgreSQLSession()
	} else {
		return initMySqlDBSession()
	}
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

func initPostgreSQLSession() *gorm.DB {
	username := os.Getenv("GORM_DATASOURCE_USERNAME")

	if username == "" {
		username = "postgres"
	}
	password := os.Getenv("GORM_DATASOURCE_PASSWORD")

	if password == "" {
		password = "postgres"
	}

	host := os.Getenv("GORM_HOST")
	if host == "" {
		host = "localhost"
	}
	dbname := os.Getenv("GORM_DBNAME")
	if dbname == "" {
		dbname = "app4"
	}
	port := os.Getenv("GORM_PORT")
	if port == "" {
		port = "5432"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, username, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	return db
}

func Save(ctx context.Context, entity DataEntity) {
	ctx, span := tracer.Start(ctx, "DataRepository#Save")
	defer span.End()

	result := repository.db.Create(&entity)

	if result.Error != nil {
		log.Error(result.Error)
	}
}
