# build stage
FROM golang:alpine AS build-env

ADD . /go/src/app

RUN set -xe \
    && apk add --update \
        git \
        bash \
    && go get -u github.com/golang/dep/cmd/dep \
    && cd /go/src/app \
    && dep ensure \
    && go build -o message-router

# final stage
FROM alpine

ENV MESSAGE_ROUTER_KAFKA_BROKERS="localhost:9092" \
    MESSAGE_ROUTER_PORT=3000

WORKDIR "/app"

COPY --from=build-env /go/src/app/message-router /app/

ENTRYPOINT [ "/app/message-router" ]
