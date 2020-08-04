package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/node"
	"github.com/joincloud/peers-touch-go/pubsub"

	_ "github.com/joincloud/peers-touch-go/logger/logrus"
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

	Sub(n.Broker())
	go func() {
		i := 1
		for {
			Pub(i, n.Broker())
			i++
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func Pub(idx int, broker pubsub.Broker) {
	// todo codec
	err := broker.Pub(context.Background(), pubsub.NewEvent(topic, pubsub.Message{
		Header: map[string]string{
			"idx": strconv.Itoa(idx),
		},
		Body: []byte("Hello, Suber"),
	}))
	if err != nil {
		panic(err)
	}
}

func Sub(broker pubsub.Broker) {
	_, err := broker.Sub(context.Background(), topic, func(event pubsub.Event) {
		logger.Infof("msg handler, topic: %s: msg: %s", topic, string(event.Message().Body))
	})
	if err != nil {
		panic(err)
	}
}
