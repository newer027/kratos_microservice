#!/bin/sh

echo "Waiting for details..."
while ! nc -z details 8081; do
    sleep 0.1
done
echo "Details started"

echo "Waiting for jaeger-collector..."
while ! nc -z jaeger-collector 9411; do
    sleep 0.1
done
echo "Jaeger-collector started"

exec "$@"
