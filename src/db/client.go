package db

import (
	"bartender/src/config"
	"fmt"

	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander/dialects/kafka"
)

var (
	ConfirmImpressionMessageBrokerClient *commander.Group
	TestMessageBrokerClient              *commander.Group
	Client                               *commander.Client
)

const (
	CONFIRM_IMPRESSION = 0
	TESTE              = 1
)

func NewConnection() {
	fmt.Println("Connecting to Kafka:", config.BrokerConnectionString)
	dialect, err := kafka.NewDialect(config.BrokerConnectionString)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	warehouse := commander.NewGroup(
		commander.NewTopic("confirm_impression", dialect, commander.CommandMessage, commander.DefaultMode),
	)
	warehouse2 := commander.NewGroup(
		commander.NewTopic("teste", dialect, commander.CommandMessage, commander.DefaultMode),
	)
	client, err := commander.NewClient(warehouse, warehouse2)
	if err != nil {
		panic(err)
	}
	// defer client.Close()
	// key, err := uuid.NewV4()
	// if err != nil {
	// panic(err)
	// }
	// command := commander.NewMessage("Available", 0, key.Bytes(), nil)
	// err = warehouse.AsyncCommand(command)
	// if err != nil {
	// fmt.Println(err)
	// }
	ConfirmImpressionMessageBrokerClient = warehouse
	TestMessageBrokerClient = warehouse2
	Client = client
}
