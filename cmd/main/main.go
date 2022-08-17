package main

import (
	"Wildberries_L0/database"
)

/*
4. Сделать валидацию данных, которые сабскрайбер получает.
6. кэш
7. http
*/
type Cache struct {
	cache map[string]int
}

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/*bytes := nats.Nats()
	var orderInfo detail.OrderInfo
	json.Unmarshal(bytes, &orderInfo)*/
}
