FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/details-linux-amd64 app
ADD ./apps/details/configs .
ADD ./scripts/wait-for .


