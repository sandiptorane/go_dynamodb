package main

import (
	"github.com/joho/godotenv"
	"go_dynamodb/database/connection"
	"go_dynamodb/handler"
	"log"
	"os"
)

func main() {
	// load env config
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("not able to load env", err)
	}

	conn, err := connection.GetConnection()
	if err != nil {
		log.Fatalln("db connect error", err)
	}

	app := handler.NewApplication(conn)

	// register routes
	r := RegisterRoutes(app)

	// run server
	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("run error:", err)
	}
}
