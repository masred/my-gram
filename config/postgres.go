package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/masred/my-gram/app/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")
	port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Photo{}, &entity.Comment{}, &entity.SocialMedia{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
