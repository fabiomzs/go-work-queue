package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	RabbitMQ RabbitMQ
}

type RabbitMQ struct {
	User               string
	Password           string
	Host               string
	VHost              string
	Port               string
	ExchangeName       string
	QueueName          string
	RoutingKey         string
	DeadLetterExchange string
	DeadLetterQueue    string
	DeadLetterTTL      int32
	RetryTime          int32
	Heartbeat          int32
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	cfg = &config{
		RabbitMQ: RabbitMQ{
			User:               viper.GetString("RABBITMQ_USER"),
			Password:           viper.GetString("RABBITMQ_PASSWORD"),
			Host:               viper.GetString("RABBITMQ_HOST"),
			VHost:              viper.GetString("RABBITMQ_VHOST"),
			Port:               viper.GetString("RABBITMQ_PORT"),
			ExchangeName:       viper.GetString("RABBITMQ_EXCHANGE"),
			QueueName:          viper.GetString("RABBITMQ_QUEUE"),
			RoutingKey:         viper.GetString("RABBITMQ_ROUTING_KEY"),
			DeadLetterExchange: viper.GetString("RABBITMQ_DEADLETTER_EXCHANGE"),
			DeadLetterQueue:    viper.GetString("RABBITMQ_DEADLETTER_QUEUE"),
			DeadLetterTTL:      viper.GetInt32("RABBITMQ_DEADLETTER_TTL"),
			RetryTime:          viper.GetInt32("RABBITMQ_RETRY_TIME"),
			Heartbeat:          viper.GetInt32("RABBITMQ_HEARTBEAT"),
		},
	}
}

func GetConfig() config {
	return *cfg
}
