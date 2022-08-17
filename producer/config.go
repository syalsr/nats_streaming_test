package main

import (
	"github.com/joho/godotenv"
	"os"
)

type natsConfig struct {
	ClusterID  string
	Channel    string
	NatsServer string
}

type natsConfigProducer struct {
	natsConfig
	ClientID string
}

func ConfigNatsStreamingProducer() *natsConfigProducer {
	err := godotenv.Load("A:/go_workspace/Wildberries_L0/.env")
	if err != nil {
		panic(err)
	}

	cfg := new(natsConfigProducer)
	cfg.NatsServer = os.Getenv("NATS_SERVER")
	cfg.ClientID = os.Getenv("CLIENTID_PRODUCER")
	cfg.Channel = os.Getenv("CHANNEL")
	cfg.ClusterID = os.Getenv("STANCLUSTERID")

	return cfg
}
