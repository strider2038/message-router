FROM golang:alpine

ENV MESSAGE_ROUTER_KAFKA_BROKERS="localhost:9092" \
    MESSAGE_ROUTER_PORT=3000

RUN set -xe \
    && apk add --update \
        git \
        bash \
    && go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/bitbucket.org/strider2038/event-router

WORKDIR "/go/src/bitbucket.org/strider2038/event-router"

RUN dep ensure \
    && go build \
    && chmod +x event-router \
    && mkdir /app \
    && cp event-router /app/event-router

WORKDIR "/app"

ENTRYPOINT [ "/app/event-router" ]
