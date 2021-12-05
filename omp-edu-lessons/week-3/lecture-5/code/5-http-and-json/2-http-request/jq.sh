#!/bin/sh

JSON=$(curl -s "https://gorest.co.in/public/v1/users?name=Agrata")

printf "%s\n" "$JSON"

echo "$JSON" | jq

echo "$JSON" | jq '.data[] | {email: .email, name: .name}'
