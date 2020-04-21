FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update && apk add bash

COPY ./dist/details-linux-amd64 /app/bin/
COPY ./build/details/run_details.sh /app/bin/
COPY ./apps/details/configs /app/bin/configs
COPY ./scripts/wait-for /app/bin/

ENV     LOG_DIR     /app/log
ENV     APP_ID      details
EXPOSE 8081
EXPOSE 9000
WORKDIR /app/bin/

ENTRYPOINT ["/app/bin/run_details.sh"]