package main

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	lib "github.com/libp2p/go-libp2p"
	ecdsa "github.com/libp2p/go-libp2p/core/crypto"
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
    // start a libp2p node with default settings
    node, err := lib.New(lib.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"))
    if err != nil {
        panic(err)
    }

    // print the node's listening addresses
    fmt.Println("Listen addresses:", node.Addrs())

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