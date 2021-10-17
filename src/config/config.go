package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var BrokerConnectionString = ""

func Load() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	BrokerConnectionString = fmt.Sprintf("brokers=%s initial-offset=newest version=%s", os.Getenv("BROKERS"), os.Getenv("CLUSTER_VERSION"))
}
