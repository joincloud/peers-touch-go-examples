module github.com/joincloud/examples

go 1.14

replace github.com/joincloud/peers-touch-go => /home/sx/Project/home/peers-touch-go

require (
	github.com/joincloud/peers-touch-go v0.0.0-20200730145552-21e6a7714fe0
	github.com/libp2p/go-libp2p v0.9.6
	github.com/libp2p/go-libp2p-autonat v0.2.3
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.5.7
	github.com/libp2p/go-libp2p-kad-dht v0.8.2
	github.com/libp2p/go-libp2p-quic-transport v0.6.0
	github.com/libp2p/go-libp2p-routing v0.0.1
	github.com/libp2p/go-libp2p-secio v0.2.2
	github.com/libp2p/go-libp2p-tls v0.1.3
)
