#!/bin/sh

echo "Waiting for redis..."
while ! nc -z redis 6379; do
    sleep 0.1
done
echo "Redis started"

echo "Waiting for memcached..."
while ! nc -z memcached 11211; do
    sleep 0.1
done
echo "Memcached started"

echo "Waiting for discovery..."
while ! nc -z discovery 7171; do
    sleep 0.1
done
echo "Discovery started"

echo "Waiting for db..."
while ! nc -z db 3306; do
    sleep 0.1
done
echo "Db started"

echo "Waiting for jaeger-collector..."
while ! nc -z jaeger-collector 9411; do
    sleep 0.1
done
echo "Jaeger-collector started"

exec "$@"
