module github.com/ozonmp/week-6-workshop

go 1.16

require (
	github.com/Shopify/sarama v1.30.0 // indirect
	github.com/ozonmp/week-6-workshop/kafka v0.0.0-20211120092448-783cfaf20bcb
)

replace github.com/ozonmp/week-6-workshop/kafka => ./kafka
