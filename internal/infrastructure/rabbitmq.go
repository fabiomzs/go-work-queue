package infrastructure

import (
	"fmt"
	"time"

	"github.com/fabiomzs/go.work-queue/configs"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQService struct {
	channel    *amqp.Channel
	connection *amqp.Connection
}

func (r *RabbitMQService) OpenConnection() {
	rabbit := configs.GetConfig().RabbitMQ

	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbit.User, rabbit.Password, rabbit.Host, rabbit.Port)
	config := amqp.Config{
		Heartbeat: time.Duration(rabbit.Heartbeat) * time.Second,
	}

	conn, err := amqp.DialConfig(amqpUrl, config)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	r.channel = ch
	r.connection = conn

	r.declareExchange()
}

func (r *RabbitMQService) CloseConnection() {
	r.channel.Close()
	r.connection.Close()
}

func NewRabbitMQService() *RabbitMQService {
	return &RabbitMQService{}
}

func (r *RabbitMQService) declareExchange() {
	rabbit := configs.GetConfig().RabbitMQ
	err := r.channel.ExchangeDeclare(
		rabbit.DeadLetterExchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	_, err = r.channel.QueueDeclare(
		rabbit.DeadLetterQueue,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    rabbit.ExchangeName,
			"x-dead-letter-routing-key": rabbit.RoutingKey,
			"x-message-ttl":             rabbit.DeadLetterTTL,
		},
	)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	r.channel.QueueBind(rabbit.DeadLetterQueue, "", rabbit.DeadLetterExchange, false, nil)

	err = r.channel.ExchangeDeclare(
		rabbit.ExchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	_, err = r.channel.QueueDeclare(
		rabbit.QueueName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange": rabbit.DeadLetterExchange,
		},
	)

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	r.channel.QueueBind(rabbit.QueueName, rabbit.RoutingKey, rabbit.ExchangeName, false, nil)
}

func (r *RabbitMQService) SendMessage(message string) error {
	rabbit := configs.GetConfig().RabbitMQ

	err := r.channel.Publish(
		rabbit.ExchangeName,
		rabbit.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: amqp.Persistent,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
