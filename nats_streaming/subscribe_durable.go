package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
	"time"
)

const clusterID = "eventbus"
const clientID = "myIdsub"

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://192.168.5.101:6222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	options := []stan.SubscriptionOption{
		stan.DurableName("my-durable"),
		stan.AckWait(20 * time.Second),
	}

	ch := make(chan stan.Msg)
	// Durable subscribe
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		ch <- *m
	}, options...)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Close()

	for {
		m := <-ch
		fmt.Printf("%v %v %s\n", m.Sequence, m.Timestamp, m.Data)
	}
}
