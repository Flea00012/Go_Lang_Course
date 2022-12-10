package proj3

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	libp2p "github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
    "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
    "github.com/libp2p/go-libp2p/core/peer"
)

type pubKey = ecdsa.PublicKey
type privKey = ecdsa.PrivateKey

type localNotifee string

func (l localNotifee) HandlePeerFound(p peer.AddrInfo)  {
    fmt.Printf("handling the found peer, with notifee: %v", l)
}

func (c ChatServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    err := r.ParseForm()
    if err != nil {
        log.Fatalln(err)
    }
}

func main() {
    sig := GetSignKeys()
    if sig == nil {
        panic("your signature was not generated")
    }

    node, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
    if err != nil {
        panic(err)
    }

    ctx := context.Background()

    var chatty ChatServer
    http.ListenAndServe(":8080", chatty)    

    privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	m := "the hash string for bytes"
	hash := sha256.Sum256([]byte(m))
	hashSign, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", hashSign)

	verified := ecdsa.VerifyASN1(&privKey.PublicKey, hash[:], hashSign)
	fmt.Println("signature verified: ", verified)

    // start a libp2p node with default settings
    host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"))
    if err != nil {
        panic(err)
    }

    // print the node's listening addresses
    fmt.Println("Listen addresses:", node.Addrs())

    pubs , err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		panic(err)
	}
    fmt.Println("the pubsub is: ‰v", pubs)

    // wait for a SIGINT or SIGTERM signal
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch
    fmt.Println("Received signal, shutting down...")

    // shut the node down
    if err := node.Close(); err != nil {
        panic(err)
    }

    var l localNotifee
   
    mdns.NewMdnsService(host, "my server name", l)
}