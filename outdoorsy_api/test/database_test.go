package database

import (
	"fmt"
	"net/http"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnect(t *testing.T) {
	var err error

	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "root"
	dbPassword := "root"
	dbName := "testingwithrentals"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser,
		dbPassword, dbName)

	_, err = gorm.Open(postgres.Open(psqlInfo))

	if err != nil {
		t.Fatalf("failed to connect database")
	}
}

func TestRouterWithParams(t *testing.T) {
	_, err := http.NewRequest(http.MethodGet, "/rentals?price_min=9000&price_max=75000", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRouter(t *testing.T) {
	_, err := http.NewRequest(http.MethodGet, "/rentals", nil)
	if err != nil {
		t.Fatal(err)
	}
}
