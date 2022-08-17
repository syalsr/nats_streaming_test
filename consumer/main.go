package main

import (
	"Wildberries_L0/cache"
	"Wildberries_L0/config"
	"Wildberries_L0/database"
	"Wildberries_L0/detail"
	"Wildberries_L0/web"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"time"
)

/*
2. Подключаемся к бд
1. Подписываем на канал натс
2. Кладем данные в бд
3. Из бд берем данные в кеш
4. html
*/

func main() {
	connection := database.Connect()
	orderCache := cache.LoadCacheFromDatabase(connection)

	cfg := config.ConfigNatsStreamingConsumer()
	sc, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsURL(cfg.NatsServer))
	if err != nil {
		panic(err)
	}

	var orders []detail.OrderInfo
	sub, err := sc.Subscribe(cfg.Channel, func(msg *stan.Msg) {
		if !json.Valid(msg.Data) {
			panic(err)
		}

		info := detail.NewOrderInfo()
		err := json.Unmarshal(msg.Data, &info)
		orderCache.SaveToCache(info)
		orders = append(orders, *info)

		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)

	database.InsertData(connection, orders)
	defer sub.Unsubscribe()

	web.Server(orderCache)
}
