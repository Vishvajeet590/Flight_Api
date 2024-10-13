package main

import (
	"Flight_Api/app/database"
	"Flight_Api/app/router"
	"log"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return
	}
	r := router.NewRouter()

	r.Run(":8080")

}
