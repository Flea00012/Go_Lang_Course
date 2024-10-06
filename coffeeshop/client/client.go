package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/fowlerlee/proto_example/coffeeshop_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("failed to connect to gRPC server")
	}

	defer conn.Close()

	coffeeshop := pb.NewCoffeeShopClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuStream, err := coffeeshop.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatal("error calling funciton getMenu")
	}

	done := make(chan bool)

	var items []*pb.Item

	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			items = resp.Items
			log.Printf("Resp received: %v", resp.Items)

		}
	}()

	<-done

	receipt, err := coffeeshop.PlaceOrder(ctx, &pb.Order{Items: items})
	if err != nil {
		log.Println("failed to place order")
	}
	log.Printf("%v", receipt)

	status, err := coffeeshop.GetOrderStatus(ctx, receipt)
	if err != nil {
		log.Println("failed to get order status")
	}
	log.Printf("%v", status)

}
