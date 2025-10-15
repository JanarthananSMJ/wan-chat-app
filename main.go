package main

import (
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

type Peer struct {
	Id   peer.ID
	Addr []multiaddr.Multiaddr
	Host host.Host
}

func CreatePeer() (*Peer, error) {
	privkey, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		return nil, fmt.Errorf("error on crypto key generation : %s", err)
	}
	host, err := libp2p.New(libp2p.Identity(privkey),
		libp2p.ListenAddrStrings("/ip6/::/tcp/0", "/ip4/0.0.0.0/tcp/0"))

	return &Peer{
		Id:   host.ID(),
		Addr: host.Addrs(),
		Host: host,
	}, nil
}

var mypeer *Peer

func main() {
	var err error
	mypeer, err := CreatePeer()
	if err != nil {
		log.Fatal("Error On Peer Creation")
	}

	fmt.Println("Peer ID : ", mypeer.Id)
	for _, addr := range mypeer.Addr {
		fmt.Printf("Peer MultiAddress : %s/p2p/%s\n", addr.String(), mypeer.Id)
	}
}
