package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)
var err error

func Connect() {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//Database connection setup
	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	//openning connection
	d, err := gorm.Open(dialect, dbURL)
	if err != nil {
		panic(err)

	} else {
		fmt.Println("Successfully connected to database")
	}
	db = d
	//Make migrations to database if they have not already been created
	//db.AutoMigrate(&Student{})
	//db.AutoMigrate(&Teacher{})
	//return db
	//defer db.Close()

}
func GetDB() *gorm.DB {
	return db
}
