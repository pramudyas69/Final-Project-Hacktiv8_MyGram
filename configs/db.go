package configs

import (
	"MyGram/models"
	"MyGram/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connection() *gorm.DB {
	config := utils.GetSQLEnv()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.Host, config.Port, config.User, config.DBName, config.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		defer log.Fatal("Connection to Database Failed")
		log.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		fmt.Println("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Photo{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
