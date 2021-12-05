package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)


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


type Consumer struct {}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var msg OrderInfo
		err := json.Unmarshal(message.Value, &msg)
		if err != nil {
			fmt.Printf("Error unmarshalling message: %s\n", err)
		}

		fmt.Printf("Msg: %v\n", msg.UserID)


		session.MarkMessage(message, "")
	}

	return nil
}


func subscribe(ctx context.Context, topic string, consumerGroup sarama.ConsumerGroup) error {
	consumer := Consumer{}

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, &consumer); err != nil {
				fmt.Printf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()


	return nil
}


var brokers  = []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}

func StartConsuming(ctx context.Context) error {

	config := sarama.NewConfig()

	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, "analytic", config)

	if err != nil {
		return err
	}


	return subscribe(ctx,"orders",  consumerGroup)
}

func main() {
	ctx := context.Background()

	err := StartConsuming(ctx)
	if err != nil {
		return
	}

	for {

	}
}