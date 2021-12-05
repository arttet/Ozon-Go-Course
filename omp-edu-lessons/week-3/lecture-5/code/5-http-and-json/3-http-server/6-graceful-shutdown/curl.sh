#!/bin/sh

while true
do
  curl localhost:8080/slow || sleep 1
done
