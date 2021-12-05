package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/ozonmp/week-6-workshop/kafka"
	"github.com/ozonmp/week-6-workshop/protos"
	"google.golang.org/protobuf/proto"
)

type DB struct {
	sync.Mutex
	Orders         map[int64]protos.Order
	OrdersStatuses map[int64]int
}

var db DB
var producer sarama.SyncProducer

const (
	OrderCreated = iota
	OrderCommited

	consumerGroup = "order_service"
)

func CreateOrder(ctx context.Context, message *sarama.ConsumerMessage) error {
	var order protos.Order

	err := proto.Unmarshal(message.Value, &order)
	if err != nil {
		return err
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.Orders[order.OrderID]; ok {
		fmt.Printf("Order with ID %d already exists\n", order.OrderID)
		return nil
	}

	db.Orders[order.OrderID] = order
	db.OrdersStatuses[order.OrderID] = OrderCreated

	if rand.Intn(10) == 1 {
		fmt.Printf("Make order ID %d as failed\n", order.OrderID)
		err = kafka.SendMessage(producer, "cancel_order", message.Value)
		if err != nil {
			return err
		}
	}

	err = kafka.SendMessage(producer, "control_products", message.Value)

	return nil
}

func CommitOrder(ctx context.Context, message *sarama.ConsumerMessage) error {
	var order protos.Order

	err := proto.Unmarshal(message.Value, &order)
	if err != nil {
		return err
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.Orders[order.OrderID]; ok {
		fmt.Printf("Order with ID %d commited", order.OrderID)
		db.OrdersStatuses[order.OrderID] = OrderCommited
		return nil
	}

	return nil
}

func CancelOrder(ctx context.Context, message *sarama.ConsumerMessage) error {
	var order protos.Order

	err := proto.Unmarshal(message.Value, &order)
	if err != nil {
		return err
	}

	db.Lock()
	defer db.Unlock()

	if _, ok := db.Orders[order.OrderID]; ok {
		fmt.Printf("Order with ID %d canceled\n", order.OrderID)
		delete(db.Orders, order.OrderID)
		return nil
	}

	return nil
}

func main() {
	brokers := []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}

	ctx := context.Background()

	db.Orders = make(map[int64]protos.Order)
	db.OrdersStatuses = make(map[int64]int)

	var err error
	producer, err = kafka.NewSyncProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	err = kafka.StartConsuming(ctx, brokers, "create_order", consumerGroup, CreateOrder)
	if err != nil {
		log.Fatal(err)
	}
	err = kafka.StartConsuming(ctx, brokers, "commit_order", consumerGroup, CommitOrder)
	if err != nil {
		log.Fatal(err)
	}
	err = kafka.StartConsuming(ctx, brokers, "cancel_order", consumerGroup, CancelOrder)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Minute * 10)
}
