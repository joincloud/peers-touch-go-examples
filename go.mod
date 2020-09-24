module github.com/joincloud/examples

go 1.14

replace (
	github.com/joincloud/peers-touch-go v1.0.0 => ../peers-touch-go
	github.com/libp2p/go-libp2p v0.9.6 => ../go-libp2p
	github.com/libp2p/go-libp2p-core v0.6.1 => ../go-libp2p-core
	github.com/ipfs/go-ipfs v0.6.0 => ../go-ipfs
	github.com/libp2p/go-libp2p-pubsub v0.3.0 => ../go-libp2p-pubsub
)

require (
	github.com/joincloud/peers-touch-go v1.0.0
	github.com/libp2p/go-libp2p v0.9.6
	github.com/libp2p/go-libp2p-autonat v0.3.2
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.6.1
	github.com/libp2p/go-libp2p-kad-dht v0.8.2
	github.com/libp2p/go-libp2p-quic-transport v0.6.0
	github.com/libp2p/go-libp2p-routing v0.0.1
	github.com/libp2p/go-libp2p-secio v0.2.2
	github.com/libp2p/go-libp2p-tls v0.1.3
)
