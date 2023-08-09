package main

import (
	"go_dynamodb/database/connection"
	"go_dynamodb/database/dynamo"
	"log"
)

func main() {
	conn, err := connection.GetConnection()
	log.Fatalln("db connect error", err)

	// get instance
	dynamo.GetInstance(conn)

}
