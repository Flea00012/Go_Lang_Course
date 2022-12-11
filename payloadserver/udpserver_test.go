package payloadserver

import (
	"fmt"
	"net"
	"context"
)

func EchoUDPServer(cxt context.Context, addr string) (net.Addr, error) {
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
