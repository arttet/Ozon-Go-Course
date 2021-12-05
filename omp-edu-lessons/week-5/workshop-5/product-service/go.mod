module github.com/ozonmp/week-5-workshop/product-service

go 1.16

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/lib/pq v1.10.3
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/grpc v1.41.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.2
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/genproto v0.0.0-20211029142109-e255c875f7c7
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Masterminds/squirrel v1.5.1
	github.com/jmoiron/sqlx v1.3.4
	github.com/ozonmp/week-5-workshop/category-service/pkg/category-service v0.0.0-20211106062202-6ef7f265314c
	github.com/pressly/goose/v3 v3.3.1
	github.com/snovichkov/zap-gelf v1.0.1
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/zap v1.13.0
)

replace github.com/ozonmp/week-5-workshop/category-service/pkg/category-service => ../category-service/pkg/category-service
