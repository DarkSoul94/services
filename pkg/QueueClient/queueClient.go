package queueclient

import "github.com/streadway/amqp"

type QueueClient interface {
	Publish(queueName string, data []byte) error
	Consume(queueName string) (<-chan amqp.Delivery, error)
}
