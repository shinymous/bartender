package advertising

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofrs/uuid"
	"github.com/jeroenrinzema/commander"
	"github.com/jeroenrinzema/commander/dialects/kafka"
	"github.com/joho/godotenv"
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

type brokerConnection struct {
	CONFIRM_IMPRESSION                   string
	REQUEST_AD                           string
	ConfirmImpressionMessageBrokerClient *commander.Group
	RequestMessageBrokerClient           *commander.Group
	Client                               *commander.Client
}

type BrokerConfig struct {
	CONFIRM_IMPRESSION string
	REQUEST_AD         string
}

type BrokerConnection interface {
	SendAsynMessage(topicName string, data interface{})
	GetBrokerConfig() BrokerConfig
}

func CreateConnection() BrokerConnection {
	var erro error
	if erro = godotenv.Load("../../.env"); erro != nil {
		log.Fatal(erro)
	}
	brokerConnectionString := fmt.Sprintf("brokers=%s initial-offset=newest version=%s", os.Getenv("BROKERS"), os.Getenv("CLUSTER_VERSION"))
	fmt.Println("Connecting to Kafka:", brokerConnectionString)
	dialect, err := kafka.NewDialect(brokerConnectionString)
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
	return &brokerConnection{
		CONFIRM_IMPRESSION:                   CONFIRM_IMPRESSION,
		REQUEST_AD:                           REQUEST_AD,
		ConfirmImpressionMessageBrokerClient: warehouse,
		RequestMessageBrokerClient:           warehouse2,
		Client:                               client,
	}
}

func (b brokerConnection) GetBrokerConfig() BrokerConfig {
	return BrokerConfig{
		CONFIRM_IMPRESSION: b.CONFIRM_IMPRESSION,
		REQUEST_AD:         b.REQUEST_AD,
	}
}

func (b brokerConnection) SendAsynMessage(topicName string, data interface{}) {
	key, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)
	fmt.Println(data)
	command := commander.NewMessage("", 0, key.Bytes(), reqBodyBytes.Bytes())
	if topicName == b.CONFIRM_IMPRESSION {
		err = ConfirmImpressionMessageBrokerClient.AsyncCommand(command)
	} else {
		err = RequestMessageBrokerClient.AsyncCommand(command)
	}
	if err != nil {
		fmt.Println(err)
	}
}
