package broker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	kafka "github.com/segmentio/kafka-go"
)

var (
	ConfirmImpressionTopicConn *kafka.Writer
	RequestAdTopicConn         *kafka.Writer
)

const (
	CONFIRM_IMPRESSION = "confirm_impression"
	REQUEST_AD         = "request_ad"
)

type brokerConnection struct {
	CONFIRM_IMPRESSION         string
	REQUEST_AD                 string
	ConfirmImpressionTopicConn *kafka.Writer
	RequestAdTopicConn         *kafka.Writer
}

type BrokerConfig struct {
	CONFIRM_IMPRESSION string
	REQUEST_AD         string
}

type BrokerConnection interface {
	SendMessage(topicName string, data interface{})
	GetBrokerConfig() BrokerConfig
}

func newKafkaWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("BROKERS")),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func createTopics() {
	conn, err := kafka.Dial("tcp", os.Getenv("BROKERS"))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             CONFIRM_IMPRESSION,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		{
			Topic:             REQUEST_AD,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func CreateConnection() BrokerConnection {
	createTopics()
	return &brokerConnection{
		CONFIRM_IMPRESSION: CONFIRM_IMPRESSION,
		REQUEST_AD:         REQUEST_AD,
	}
}

func (b brokerConnection) GetBrokerConfig() BrokerConfig {
	return BrokerConfig{
		CONFIRM_IMPRESSION: b.CONFIRM_IMPRESSION,
		REQUEST_AD:         b.REQUEST_AD,
	}
}

func (b brokerConnection) SendMessage(topicName string, data interface{}) {
	writer := newKafkaWriter(topicName)
	defer writer.Close()
	key, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)
	msg := kafka.Message{
		Key:   []byte(key.String()),
		Value: reqBodyBytes.Bytes(),
		Time:  time.Now().UTC(),
	}
	if topicName == b.CONFIRM_IMPRESSION {
		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			fmt.Println(err.Error())
		}
	}
}
