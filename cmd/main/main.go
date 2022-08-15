package main

import (
	"fmt"
	sql "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "syalsr" +
		""
	password = "12345"
	dbname   = "wildberries_l0"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open(psqlconn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connected!")
}

func nats() {
	/*sc, _ := stan.Connect(clusterID, clientID)

	// Simple Synchronous Publisher
	sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	sc.Close()*/
}
