FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update && apk add bash

COPY ./dist/reviews-linux-amd64 /app/bin/
COPY ./build/reviews/run_reviews.sh /app/bin/
COPY ./apps/reviews/configs /app/bin/configs
COPY ./scripts/wait-for /app/bin/

ENV     LOG_DIR     /app/log
ENV     APP_ID      reviews
EXPOSE 8083
EXPOSE 9002
WORKDIR /app/bin/

ENTRYPOINT ["/app/bin/run_reviews.sh"]