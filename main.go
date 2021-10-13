package main

import (
	"bartender/src/router"
	"fmt"
	"log"
)

func main() {
	app := router.Generate()
	port := "5000"
	log.Fatal(app.Listen(fmt.Sprintf("127.0.0.1:%s", port)))
}
