/**
	Rahmat wahyu hadi
**/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./application"
	"./config"
	"./models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func InitDatabaseConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(":: Checking databse connection ..... ")
	db, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatal("!!! Could not connect to the database")
	}

	fmt.Println(":: Database Connected")
	defer db.Close()
}

func initApps() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(":: APP running on port " + os.Getenv("API_PORT"))
}

func InitRouter() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := application.NewServer()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), server))
}

func MigrateDatabase() {
	fmt.Println(":: Migration Databases .....")
	db := config.GetDatabaseConnection()
	db.AutoMigrate(&models.User{})
	fmt.Println(":: Migration Databases Done")

}

// Main functions dan initialdata
func main() {
	InitDatabaseConnection()
	MigrateDatabase()
	initApps()
	InitRouter()
}
