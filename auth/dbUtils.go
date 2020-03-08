package auth

import (
	"fmt"
	// "os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
)


// ConnectToPostgres connects to Postgres Database
func ConnectToPostgres() *gorm.DB {
	// username := os.Getenv("databaseUser")
	// password := os.Getenv("databasePassword")
	// databaseName := os.Getenv("databaseName")
	// databaseHost := os.Getenv("databaseHost")

	username := "florian"
	password := "test"
	databaseName := "podcasts"
	databaseHost := "postgres"
	port := "5111"

	//Define DB connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s, sslmode=disable", databaseHost, port, username, databaseName, password)

	//connect to db URI
	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	// close db when not in use
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(
		&User{})

	fmt.Println("Successfully connected!", db)
	return db

}
