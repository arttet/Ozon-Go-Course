module github.com/ozonmp/week-6-workshop/warehouse

go 1.16

require (
	github.com/Shopify/sarama v1.30.0
	github.com/ozonmp/week-6-workshop/kafka v0.0.0-20211120092448-783cfaf20bcb
	github.com/ozonmp/week-6-workshop/protos v0.0.0-20211120103414-1d61439690d5
	google.golang.org/protobuf v1.27.1
)

replace github.com/ozonmp/week-6-workshop/kafka => ../kafka
