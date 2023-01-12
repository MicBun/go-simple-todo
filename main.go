package main

import (
	"github.com/MicBun/go-simple-todo/config"
	"github.com/MicBun/go-simple-todo/docs"
	"github.com/MicBun/go-simple-todo/route"
	"github.com/joho/godotenv"
	"log"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default env")
	}

	description := "This is a simple todo list.\n\n" +
		"To get Bearer Token, first you need to register then login.\n\n" +
		"Checkout my Github: https://github.com/MicBun\n\n" +
		"Checkout my Linkedin: https://www.linkedin.com/in/MicBun\n\n"

	docs.SwaggerInfo.Title = "Activity Tracking API"
	docs.SwaggerInfo.Description = description

	database := config.ConnectDataBase()
	sqlDB, _ := database.DB()
	defer sqlDB.Close()
	r := route.SetupRouter(database)
	r.Run()
}
