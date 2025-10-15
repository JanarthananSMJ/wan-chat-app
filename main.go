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
	Id        peer.ID
	Addr      []multiaddr.Multiaddr
	Host      host.Host
	Bootstrap []peer.AddrInfo
}

var BootstrapAddresses = []string{
	"/ip4/147.75.83.83/tcp/4001/p2p/QmNnooDu7fKXGG5kZsXK2xhmijT2v6d7HJv9Vid8yymvCj",
	"/ip6/2604:1380:1000:6000::1/tcp/4001/p2p/QmNnooDu7fKXGG5kZsXK2xhmijT2v6d7HJv9Vid8yymvCj",
	"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	"/ip6/2604:a880:1:20::203:d001/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
}

func CreatePeer() (*Peer, error) {
	privkey, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		return nil, fmt.Errorf("error on crypto key generation : %s", err)
	}
	host, err := libp2p.New(libp2p.Identity(privkey),
		libp2p.ListenAddrStrings("/ip6/::/tcp/0", "/ip4/0.0.0.0/tcp/0"))

	var bootstraps []peer.AddrInfo
	for _, addrStr := range BootstrapAddresses {
		maddr, err := multiaddr.NewMultiaddr(addrStr)
		if err != nil {
			fmt.Println("Error on Bootstrap Address : ", err)
		}
		info, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			fmt.Println("Failed to parse bootstrap info:", err)
			continue
		}
		bootstraps = append(bootstraps, *info)
	}
	return &Peer{
		Id:        host.ID(),
		Addr:      host.Addrs(),
		Host:      host,
		Bootstrap: bootstraps,
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

	fmt.Println("\t\tBootstrap Peers\t\t")
	for _, b := range mypeer.Bootstrap {
		for _, a := range b.Addrs {
			fmt.Printf("Bootstrap : %s/p2p/%s\n", a, b.ID)
		}
	}
}
