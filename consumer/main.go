package main

import (
	"Wildberries_L0/cache"
	"Wildberries_L0/detail"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("A:/go_workspace/Wildberries_L0/.env")
	clusterID := os.Getenv("STANCLUSTERID")
	clientID := os.Getenv("CLIENTID_CONSUMER")
	channel := os.Getenv("CHANNEL")
	natsServer := os.Getenv("NATS_SERVER")

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsServer))
	if err != nil {
		panic(err)
	}
	order_cache := cache.NewCache()
	sub, err := sc.Subscribe(channel, func(msg *stan.Msg) {
		if !json.Valid(msg.Data) {
			panic(err)
		}

		var info detail.OrderInfo
		err := json.Unmarshal(msg.Data, &info)
		if err != nil {
			panic(err)
		}

		order_cache.SaveToCache(&info)
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(18 * time.Second)
	for key, value := range order_cache.Getc() { //перебираем всю мапу
		fmt.Println(key, value)
	}
	defer sub.Unsubscribe()
}
