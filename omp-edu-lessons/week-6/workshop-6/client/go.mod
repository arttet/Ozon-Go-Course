module github.com/ozonmp/week-6-workshop/orders

go 1.16

replace (
	github.com/ozonmp/week-6-workshop/kafka => ../kafka
	github.com/ozonmp/week-6-workshop/protos => ../protos
)

require (
	github.com/ozonmp/week-6-workshop/kafka v0.0.0-00010101000000-000000000000
	github.com/ozonmp/week-6-workshop/protos v0.0.0-20211120092448-783cfaf20bcb
	google.golang.org/protobuf v1.27.1
)
