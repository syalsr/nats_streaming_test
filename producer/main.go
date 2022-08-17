package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"os"
)

func main() {
	err := godotenv.Load("A:/go_workspace/Wildberries_L0/.env")
	clusterID := os.Getenv("STANCLUSTERID")
	clientID := os.Getenv("CLIENTID_PRODUCER")
	channel := os.Getenv("CHANNEL")
	natsServer := os.Getenv("NATS_SERVER")
	fmt.Println(natsServer, clusterID, clientID)
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsServer))
	defer sc.Close()
	if err != nil {
		panic(err)
	}

	for {
		//A:/go_workspace/Wildberries_L0/model.json
		var file_path string
		fmt.Println("Enter path to file")
		fmt.Scan(&file_path)
		content, err := os.ReadFile(file_path)
		if err != nil {
			fmt.Println("Cannot read a file, may be incorrect path")
			continue
		}
		err = sc.Publish(channel, content)
		if err != nil {
			panic(err)
		}
	}
}
