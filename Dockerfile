FROM golang:alpine

ENV MESSAGE_ROUTER_KAFKA_BROKERS="localhost:9092" \
    MESSAGE_ROUTER_PORT=3000

RUN set -xe \
    && apk add --update \
        git \
        bash \
    && go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/strider2038/message-router

WORKDIR "/go/src/github.com/strider2038/message-router"

RUN dep ensure \
    && go build \
    && chmod +x message-router \
    && mkdir /app \
    && cp message-router /app/message-router

WORKDIR "/app"

ENTRYPOINT [ "/app/message-router" ]
