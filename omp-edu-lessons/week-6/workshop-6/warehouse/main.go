package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"

	"github.com/Shopify/sarama"
	"github.com/ozonmp/week-6-workshop/kafka"
	"github.com/ozonmp/week-6-workshop/protos"
	"google.golang.org/protobuf/proto"
)

const (
	controlProducts = "control_products"
	cancelOrder     = "cancel_order"
	commitOrder     = "commit_order"

	consumerGroup = "warehouse_service"
)

var producer sarama.SyncProducer

func checkAndReserv(product *protos.Product) error {
	if rand.Intn(10) == 1 {
		fmt.Printf("%v is out of stock\n", product.SKU)
		return errors.New("out of stock")
	}
	return nil
}

func handleCancelOrders(ctx context.Context, message *sarama.ConsumerMessage) error {
	var order protos.Order

	err := proto.Unmarshal(message.Value, &order)
	if err != nil {
		fmt.Printf("can't unmarshall: %v\n", err)
		return err
	}
	fmt.Printf("cancle order %v\n", order.OrderID)
	return nil
}

func handleReserveOrders(ctx context.Context, message *sarama.ConsumerMessage) error {
	var order protos.Order

	err := proto.Unmarshal(message.Value, &order)
	if err != nil {
		fmt.Printf("unmarshall error: %v\n", err)
		return err
	}

	orderOk := true
	for _, item := range order.Products {
		if checkAndReserv(item) != nil {
			orderOk = false
			break
		}
	}

	if orderOk {
		fmt.Printf("order %v is ok\n", order.OrderID)
		err = kafka.SendMessage(producer, commitOrder, message.Value)
	} else {
		fmt.Printf("order %v is out of stock\n", order.OrderID)
		err = kafka.SendMessage(producer, cancelOrder, message.Value)
	}
	if err != nil {
		fmt.Printf("on produce: %v\n", err)
		return err
	}

	return nil
}

func main() {
	brokers := []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var err error

	producer, err = kafka.NewSyncProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	err = kafka.StartConsuming(ctx, brokers, controlProducts, consumerGroup, handleReserveOrders)
	if err != nil {
		log.Fatal(err)
	}

	err = kafka.StartConsuming(ctx, brokers, cancelOrder, consumerGroup, handleCancelOrders)
	if err != nil {
		log.Fatal(err)
	}

	for {
	}
}
