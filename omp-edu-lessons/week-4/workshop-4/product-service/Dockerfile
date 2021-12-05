# Builder

FROM golang:1.16-alpine AS builder
RUN apk add --update make git protoc protobuf protobuf-dev curl gcc musl-dev

ARG GITHUB_PATH=github.com/ozonmp/week-4-workshop/product-service

COPY Makefile /home/${GITHUB_PATH}/Makefile
COPY go.mod /home/${GITHUB_PATH}/go.mod 
COPY go.sum /home/${GITHUB_PATH}/go.sum
COPY pkg /home/${GITHUB_PATH}/pkg

WORKDIR /home/${GITHUB_PATH}
RUN make deps

COPY . /home/${GITHUB_PATH}
RUN make build

# gRPC Server

FROM alpine:latest as server
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

ARG GITHUB_PATH=github.com/ozonmp/week-4-workshop/product-service

COPY --from=builder /home/${GITHUB_PATH}/bin/grpc-server .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
COPY --from=builder /home/${GITHUB_PATH}/swagger ./swagger
COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root grpc-server

EXPOSE 50051
EXPOSE 8080
EXPOSE 9100

CMD ["./grpc-server"]
