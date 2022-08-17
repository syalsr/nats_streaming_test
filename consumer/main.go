package main

import (
	"Wildberries_L0/config"
	"Wildberries_L0/database"
	"Wildberries_L0/detail"
	"encoding/json"
	"fmt"
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
	cfg := config.ConfigNatsStreamingConsumer()
	sc, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsURL(cfg.NatsServer))
	if err != nil {
		panic(err)
	}
	var info detail.OrderInfo
	sub, err := sc.Subscribe(cfg.Channel, func(msg *stan.Msg) {
		if !json.Valid(msg.Data) {
			panic(err)
		}

		err := json.Unmarshal(msg.Data, &info)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(18 * time.Second)
	defer sub.Unsubscribe()

	var info detail.OrderInfo
	json.Unmarshal([]byte(js), &info)
	connection := database.Connect()
	//database.InsertData(connection, info)
	info1 := database.GetUID(connection, info.OrderUID)
	fmt.Println(info1)
}
