package main

import (
	"context"

	"github.com/joincloud/peers-touch-go/node"
)

func main() {
	ctx := context.Background()
	n, err := node.NewNode(ctx)
	if err != nil {
		panic(err)
	}

	n.Broker()
}
