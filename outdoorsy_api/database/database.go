package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("dbHost"), os.Getenv("dbPort"), os.Getenv("dbUser"),
		os.Getenv("dbPassword"), os.Getenv("dbName"))

	DB, err = gorm.Open(postgres.Open(psqlInfo))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}
