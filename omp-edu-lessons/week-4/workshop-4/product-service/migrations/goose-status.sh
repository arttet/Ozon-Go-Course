#!/bin/sh

goose -dir migrations \
  postgres "user=user password=password host=localhost port=5432 database=db sslmode=disable" \
  status