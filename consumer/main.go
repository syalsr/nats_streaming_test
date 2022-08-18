package main

import (
	"Wildberries_L0/cache"
	"Wildberries_L0/config"
	"Wildberries_L0/database"
	"Wildberries_L0/detail"
	"Wildberries_L0/web"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

func main() {
	fmt.Println("Connect to database")
	connection := database.Connect()

	fmt.Println("Load data from database to cache")
	orderCache := cache.LoadCacheFromDatabase(connection)

	cfg := config.ConfigNatsStreamingConsumer()
	sc, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsURL(cfg.NatsServer))
	if err != nil {
		panic(err)
	}

	fmt.Println("Subscribing")
	var orders []detail.OrderInfo
	sub, err := sc.Subscribe(cfg.Channel, func(msg *stan.Msg) {
		if !json.Valid(msg.Data) {
			panic(err)
		}

		info := detail.NewOrderInfo()
		err := json.Unmarshal(msg.Data, &info)
		if err != nil {
			panic(err)
		}
		orderCache.SaveToCache(info)
		orders = append(orders, *info)
	})
	if err != nil {
		orderCache.SendFromCacheToDatabase(connection)
		panic(err)
	}
	time.Sleep(5 * time.Second)

	database.InsertData(connection, orders)
	defer sub.Unsubscribe()

	fmt.Println("Starting server: http://localhost:8080")
	web.Server(orderCache)
}
