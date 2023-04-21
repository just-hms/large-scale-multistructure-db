#!/bin/sh

echo "Waiting for db..."

while ! nc -z db 27017; do
    sleep 0.1
done

echo "db started"


echo "Waiting for cache..."

while ! nc -z cache 6379; do
    sleep 0.1
done

echo "cache started"

go run ./...