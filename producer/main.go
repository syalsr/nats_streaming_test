package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"os"
	"path/filepath"
)

func main() {

	cfg := ConfigNatsStreamingProducer()
	sc, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsURL(cfg.NatsServer))
	defer sc.Close()
	if err != nil {
		panic(err)
	}

	//A:/go_workspace/Wildberries_L0/models
	var dir string
	for {
		fmt.Println("Enter path to folder")
		fmt.Scan(&dir)

		files, err := os.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
				content, err := os.ReadFile(dir + "/" + file.Name())
				if err != nil {
					fmt.Println("Cannot read a file, may be incorrect path")
					continue
				}
				err = sc.Publish(cfg.Channel, content)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
