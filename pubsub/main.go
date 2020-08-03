package main

import (
	"context"
	"fmt"
	"github.com/joincloud/peers-touch-go/pubsub"
	"time"

	"github.com/joincloud/peers-touch-go/node"
)

var (
	topic = "touch-go.pubsub.example"
)

func main() {
	ctx := context.Background()
	n, err := node.NewNode(ctx)
	if err != nil {
		panic(err)
	}

	go Sub(n.Broker())
	go func() {
		for {
			Pub(n.Broker())
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func Pub(broker pubsub.Broker) {
	err := broker.Pub(context.Background(), pubsub.NewEvent(topic, []byte("Hello, Suber")))
	if err != nil {
		panic(err)
	}
}

func Sub(broker pubsub.Broker) {
	_, err := broker.Sub(context.Background(), topic, func(event pubsub.Event) {
		fmt.Printf("%s: %s\n", topic, event.Message())
	})
	if err != nil {
		panic(err)
	}
}
