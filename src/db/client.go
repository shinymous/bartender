package db

import (
	"bartender/src/config"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander/dialects/kafka"
)

var (
	ConfirmImpressionMessageBrokerClient *commander.Group
	RequestMessageBrokerClient           *commander.Group
	Client                               *commander.Client
)

const (
	CONFIRM_IMPRESSION = "confirm_impression"
	REQUEST_AD         = "request_ad"
)

func NewConnection() {
	fmt.Println("Connecting to Kafka:", config.BrokerConnectionString)
	dialect, err := kafka.NewDialect(config.BrokerConnectionString)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	warehouse := commander.NewGroup(
		commander.NewTopic(CONFIRM_IMPRESSION, dialect, commander.CommandMessage, commander.DefaultMode),
	)
	warehouse2 := commander.NewGroup(
		commander.NewTopic(REQUEST_AD, dialect, commander.CommandMessage, commander.DefaultMode),
	)
	client, err := commander.NewClient(warehouse, warehouse2)
	if err != nil {
		panic(err)
	}
	ConfirmImpressionMessageBrokerClient = warehouse
	RequestMessageBrokerClient = warehouse2
	Client = client
}

func SendAsynMessage(topicName string, data interface{}) {
	key, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)
	fmt.Println(data)
	command := commander.NewMessage("", 0, key.Bytes(), reqBodyBytes.Bytes())
	if topicName == CONFIRM_IMPRESSION {
		err = ConfirmImpressionMessageBrokerClient.AsyncCommand(command)
	} else {
		err = RequestMessageBrokerClient.AsyncCommand(command)
	}
	if err != nil {
		fmt.Println(err)
	}
}
