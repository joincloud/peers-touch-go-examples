package main

import (
	"context"
	"fmt"
	ma "github.com/multiformats/go-multiaddr"
	"strconv"
	"time"

	"github.com/joincloud/peers-touch-go/logger"
	"github.com/joincloud/peers-touch-go/node"
	"github.com/joincloud/peers-touch-go/pubsub"
	"github.com/libp2p/go-libp2p"
	// use json codec
	_ "github.com/joincloud/peers-touch-go/codec/json"
	_ "github.com/joincloud/peers-touch-go/logger/logrus"
)

var (
	topic = "touch-go.pubsub.example"
)

func main() {
	// The context governs the lifetime of the libp2p node.
	// Cancelling it will stop the the host.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	listen, _ := ma.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", 8899))
	host, _ := libp2p.New(ctx, libp2p.ListenAddrs(listen))
	n, err := node.NewNode(ctx, node.Host(host))
	if err != nil {
		panic(err)
	}

	logger.Infof("id: %s", n.ID())

	_, err = n.Broker().Sub(ctx, topic, func(event pubsub.Event) {
		logger.Infof("msg handler, topic: %s: msg: %s", topic, event.Message().Body)
	})
	if err != nil {
		panic(err)
	}

	go func() {
		i := 1
		for {
			err := n.Broker().Pub(ctx, pubsub.NewEvent(topic, pubsub.Message{
				Header: map[string]string{
					"idx": strconv.Itoa(i),
				},
				Body: "Hello, Suber",
			}))
			if err != nil {
				panic(err)
			}
			i++
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-time.After(300 * time.Second):
		fmt.Println("timeout 2")
	}
}
