FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update && apk add bash

COPY ./dist/products-linux-amd64 /app/bin/
COPY ./build/products/run_products.sh /app/bin/
COPY ./apps/products/configs /app/bin/configs
COPY ./scripts/wait-for /app/bin/

ENV     LOG_DIR     /app/log
ENV     APP_ID      products
EXPOSE 8080
WORKDIR /app/bin/

ENTRYPOINT ["/app/bin/run_products.sh"]

