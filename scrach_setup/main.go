package main

import (
	"log"
	"scrach_setup/Config"
	"scrach_setup/Migrations"

	"scrach_setup/Routes"

	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	////initialize the database
	Config.InitDB()
	///add this line to main.go to initialize the migration
	Migrations.Migrate()

}
func main() {

	//setup routes
	r := Routes.SetupRouter()

	// running
	r.Run(":8080")
}
