#!/bin/sh

GRPC_HOST="localhost:6002"
GRPC_METHOD="ozonmp.week_4_workshop.category_service.category_service.v1.CategoryService/GetCategoryById"

payload=$(
  cat <<EOF
{
  "id": 4
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -rpc-header 'x-app-name:dev' \
  -rpc-header 'x-app-version:1' \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}
