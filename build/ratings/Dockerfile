FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update && apk add bash

COPY ./dist/ratings-linux-amd64 /app/bin/
COPY ./build/ratings/run_ratings.sh /app/bin/
COPY ./apps/ratings/configs /app/bin/configs
COPY ./scripts/wait-for /app/bin/

ENV     LOG_DIR     /app/log
ENV     APP_ID      ratings
EXPOSE 8082
EXPOSE 9001
WORKDIR /app/bin/

ENTRYPOINT ["/app/bin/run_ratings.sh"]