package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Shopify/sarama"
)

var brokers = []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}

func newSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func newAsyncProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(brokers, config)

	return producer, err
}

func prepareMessage(topic string, message []byte) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(message),
	}

	return msg
}

type ProductInfo struct {
	SKU 	int64
	Price	float64
	Cnt		int64
}

type OrderInfo struct {
	UserID 		int64
	CreatedAt	time.Time
	Products    []ProductInfo
}

func generateOrder() OrderInfo {
	productCount := rand.Intn(20)

	products := make([]ProductInfo, 0, productCount)
	for i := 0; i < productCount; i++ {
		products = append(products, ProductInfo{
			SKU:   int64(rand.Intn(100)),
			Price: float64(rand.Intn(100)) / 100,
			Cnt:   int64(rand.Intn(10)),
		})
	}
	oi := OrderInfo{
		UserID: int64(rand.Intn(100)),
		CreatedAt: time.Now().UTC(),
		Products: products,
	}

	return oi
}

func main() {

	syncProducer, err := newSyncProducer()
	if err != nil {
		log.Fatal(err)
	}

	asyncProducer, err := newAsyncProducer()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for err := range asyncProducer.Errors() {
			fmt.Printf("Msg async err: %v\n", err)
		}
	 }()

	go func() {
		for succ := range asyncProducer.Successes() {
			fmt.Printf("Msg written async. Patrition: %d. Ossfet: %d\n", succ.Partition, succ.Offset)
		}
	}()

	for {
		oi := generateOrder()

		oiJson, err := json.Marshal(&oi)
		if err != nil {
			log.Fatal(err)
		}

		msg := prepareMessage("orders", oiJson)

		if rand.Int() % 2 == 0 {
			partition, offset, err := syncProducer.SendMessage(msg)
			if err != nil {
				fmt.Printf("Msg sync err: %v\n", err)
			} else {
				fmt.Printf("Msg written sync. Patrition: %d. Ossfet: %d\n", partition, offset)
			}
		} else {
			asyncProducer.Input() <- msg
		}

		time.Sleep(500*time.Millisecond)
	}

}