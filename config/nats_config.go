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

func ConfigNatsStreamingConsumer() *natsConfigConsumer {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	cfg := new(natsConfigConsumer)
	cfg.NatsServer = os.Getenv("NATS_SERVER")
	cfg.ClientID = os.Getenv("CLIENTID_CONSUMER")
	cfg.Channel = os.Getenv("CHANNEL")
	cfg.ClusterID = os.Getenv("STANCLUSTERID")

	return cfg
}
