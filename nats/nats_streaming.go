package nats

import (
	"github.com/nats-io/stan.go"
)

func Nats() []byte {
	sc, _ := stan.Connect("test-cluster", "simple-pub")
	defer sc.Close()

	var sub stan.Subscription
	defer sub.Unsubscribe()
	var result []byte
	sub, _ = sc.Subscribe("foo", func(m *stan.Msg) {
		result = m.Data
	})
	return result
}
