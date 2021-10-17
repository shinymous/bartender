package main

import (
	"bartender/src/config"
	"bartender/src/db"
	"bartender/src/router"
	"fmt"
	"log"
)

func main() {
	app := router.Generate()
	port := "5000"
	config.Load()
	db.NewConnection()
	log.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", port)))
}
