package main

import (
	"context"
	"example/test/chatroom"
	"fmt"
	"net"
)

func UDPServer(cxt context.Context, addr string) (net.Addr, error) {
	serv, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("bind to udp %s; %w", addr, err)
	}
	fmt.Println("successful bind to udp ...")

	go func() {
		go func() {
			<-cxt.Done()
			_ = serv.Close()
		}()

		buffer := make([]byte, 1024)

		for {
			n, clientAddr, err := serv.ReadFrom(buffer)
			if err != nil {
				return
			}
			_, err = serv.WriteTo(buffer[:n], clientAddr)

			if err != nil {
				return
			}
		}
		
	}()
	defer serv.Close()
	return serv.LocalAddr(), nil
}

func main() {
	message := chatroom.ChatMessage{
		Message:    "hej kan du ringa mig",
		SenderID:   "minIDkortOchNummer",
		SenderNick: "mitt name",
	}
	fmt.Printf("my message: %s", message)
	UDPServer(context.Background(), "/ip4/127.0.0.1/tcp/2000")
}
