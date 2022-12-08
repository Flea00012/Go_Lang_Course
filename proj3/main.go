package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
	"os/signal"
	"syscall"
    pubsub "github.com/libp2p/go-libp2p-pubsub"
	lib "github.com/libp2p/go-libp2p"
)

type pubKey = ecdsa.PublicKey
type privKey = ecdsa.PrivateKey


type ChatRoom struct {
    roomName string
    Message chan *ChatMessage
}

type ChatMessage struct {
    Message string
    SenderID string
    SenderNick string
}

func main() {

    ctx := context.Background()

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
    node, err := lib.New(lib.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"))
    if err != nil {
        panic(err)
    }

    // print the node's listening addresses
    fmt.Println("Listen addresses:", node.Addrs())

    ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}

    // wait for a SIGINT or SIGTERM signal
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch
    fmt.Println("Received signal, shutting down...")

    // shut the node down
    if err := node.Close(); err != nil {
        panic(err)
    }
   
}