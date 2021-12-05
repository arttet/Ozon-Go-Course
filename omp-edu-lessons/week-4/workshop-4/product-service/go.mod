module github.com/ozonmp/week-4-workshop/product-service

go 1.16

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/lib/pq v1.10.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rs/zerolog v1.24.0
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
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/Masterminds/squirrel v1.5.1
	github.com/jmoiron/sqlx v1.3.4
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/ozonmp/week-4-workshop/category-service/pkg/category-service v0.0.0-20211106062202-6ef7f265314c
	github.com/pressly/goose/v3 v3.3.1
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

// replace github.com/ozonmp/week-4-workshop/product-service/pkg/product-service => ./pkg/product-service
