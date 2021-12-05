# Builder

ARG GITHUB_PATH=github.com/ozonmp/week-3-workshop/product-service

FROM golang:1.16-alpine AS builder
RUN apk add --update make git protoc protobuf protobuf-dev curl
COPY . /home/${GITHUB_PATH}
WORKDIR /home/${GITHUB_PATH}
RUN make deps && make build

# gRPC Server

FROM alpine:latest as server
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/grpc-server .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
COPY --from=builder /home/${GITHUB_PATH}/migrations/ ./migrations

RUN chown root:root grpc-server

EXPOSE 50051
EXPOSE 8080
EXPOSE 9100

CMD ["./grpc-server"]
