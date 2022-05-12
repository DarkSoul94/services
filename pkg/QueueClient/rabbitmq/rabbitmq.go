package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Connect(login, pass, host string, port int) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", login, pass, host, port))
	if err != nil {
		return nil, nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, channel, nil
}

type RabbitClient struct {
	channel *amqp.Channel
}

func NewRabbitClient(ch *amqp.Channel) *RabbitClient {
	return &RabbitClient{
		channel: ch,
	}
}

func (c *RabbitClient) Publish(queueName string, data []byte) error {
	queue, err := c.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return c.channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         data,
	})
}

func (c *RabbitClient) Consume(queueName string) (<-chan amqp.Delivery, error) {
	return c.channel.Consume(queueName, "", false, false, false, false, nil)
}
