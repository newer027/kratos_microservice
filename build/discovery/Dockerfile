FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update && apk add bash

COPY ./dist/discovery-linux-amd64 /app/bin/
COPY ./apps/discovery/cmd/discovery/discovery-example.toml /app/conf/

# Copy from local build
#ADD  dist/ /app/
ENV  LOG_DIR    /app/log
EXPOSE 7171
WORKDIR /app/
CMD /app/bin/discovery-linux-amd64 -conf /app/conf/ -confkey discovery-example.toml
