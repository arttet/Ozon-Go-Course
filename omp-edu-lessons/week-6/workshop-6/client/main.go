package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ozonmp/week-6-workshop/kafka"
	"github.com/ozonmp/week-6-workshop/protos"
	"google.golang.org/protobuf/proto"
)

func generateOrder() *protos.Order {
	productCount := rand.Intn(20)

	products := make([]*protos.Product, 0, productCount)
	for i := 0; i < productCount; i++ {
		products = append(products, &protos.Product{
			SKU:   int64(rand.Intn(100)),
			Price: float64(rand.Intn(100)) / 100,
			Cnt:   int64(rand.Intn(10)),
		})
	}
	oi := protos.Order{
		UserID:    int64(rand.Intn(100)),
		Timestamp: 0,
		Products:  products,
		OrderID:   int64(rand.Intn(1000)),
	}

	return &oi
}


func main() {
	brokers := []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}

	syncProducer, err := kafka.NewSyncProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}



	for {
		oi := generateOrder()

		msg, err := proto.Marshal(oi)
		if err != nil {
			log.Fatal(err)
		}

		err = kafka.SendMessage(syncProducer, "create_order", msg)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(500*time.Millisecond)
	}
}