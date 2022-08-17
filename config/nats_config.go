package config

import (
	"github.com/joho/godotenv"
	"os"
)

type natsConfig struct {
	ClusterID  string
	Channel    string
	NatsServer string
}
type natsConfigConsumer struct {
	natsConfig
	ClientID string
}
type natsConfigProducer struct {
	natsConfig
	clientID string
}

func ConfigNatsStreamingConsumer() *natsConfigConsumer {
	err := godotenv.Load("A:/go_workspace/Wildberries_L0/.env")
	if err != nil {
		panic(err)
	}

	cfg := new(natsConfigConsumer)
	cfg.natsServer = os.Getenv("NATS_SERVER")
	cfg.clientID = os.Getenv("CLIENTID_CONSUMER")
	cfg.channel = os.Getenv("CHANNEL")
	cfg.clusterID = os.Getenv("STANCLUSTERID")

	return cfg
}

func ConfigNatsStreamingProducer() *natsConfigProducer {
	err := godotenv.Load("A:/go_workspace/Wildberries_L0/.env")
	if err != nil {
		panic(err)
	}

	cfg := new(natsConfigProducer)
	cfg.natsServer = os.Getenv("NATS_SERVER")
	cfg.clientID = os.Getenv("CLIENTID_CONSUMER_PRODUCER")
	cfg.channel = os.Getenv("CHANNEL")
	cfg.clusterID = os.Getenv("STANCLUSTERID")

	return cfg
}
