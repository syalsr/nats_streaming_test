package main

import (
	"Wildberries_L0/database"
	"Wildberries_L0/detail"
	"Wildberries_L0/nats"
	"encoding/json"
	"fmt"
)

/*
1. Сначала надо разобраться с натс стриминг. Что такое, как работает. И отправить данные из одного мейна в другой.
3. Сделать паблишер, который будет отправлять данные через натс и проверить, что сабскрайбер получает эти данные.
4. Сделать валидацию данных, которые сабскрайбер получает.
5. Потом загуглить, как подрубать базу данных и оправить данные из структуры в бд.
6. кэш
7. http
*/
type Cache struct {
	cache map[string]int
}

func main() {
	s := nats.Nats()
	fmt.Println(string(s))

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bytes := nats.Nats()
	var orderInfo detail.OrderInfo
	json.Unmarshal(bytes, &orderInfo)
}
