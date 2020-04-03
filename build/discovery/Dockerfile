FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/discovery-linux-amd64 app
ADD ./scripts/discovery/cmd/discovery-example.toml .


