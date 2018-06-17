FROM golang:alpine

RUN set -xe \
    && apk add --update \
        git \
        bash \
    && go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/bitbucket.org/strider2038/event-router

WORKDIR "/go/src/bitbucket.org/strider2038/event-router"

RUN dep ensure \
    && go build

ENTRYPOINT "event-router"
