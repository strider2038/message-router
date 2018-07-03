# Message router

[![Build Status](https://travis-ci.org/strider2038/message-router.svg?branch=master)](https://travis-ci.org/strider2038/message-router)
[![Go Report Card](https://goreportcard.com/badge/github.com/strider2038/message-router)](https://goreportcard.com/report/github.com/strider2038/message-router)
[![codecov](https://codecov.io/gh/strider2038/message-router/branch/master/graph/badge.svg)](https://codecov.io/gh/strider2038/message-router)

Microservice for dispatching messages received by HTTP (JSON RPC v2) to Apache Kafka.

## How to use

### Docker

Pull and run docker image. Don't forget to set Apache Kafka brokers host by `MESSAGE_ROUTER_KAFKA_BROKERS` environment variable.

```bash
docker pull strider2038/message-router
docker run --name message-router -d -p 3000:3000 -e "MESSAGE_ROUTER_KAFKA_BROKERS: http://kafka.com:9092" strider2038/message-router
```

Now you can send messages that will be dispatched to Apache Kafka.

```http
POST http://localhost:3000/rpc
Content-Type: application/json

[
    {
        "jsonrpc": "2.0",
        "method": "dispatch",
        "params": {
            "topic": "firstTopic",
            "message": {
                "value": {
                    "property": "value",
                    "object": {
                        "a": 1
                    }
                }
            }
        },
        "id": 1
    },
    {
        "jsonrpc": "2.0",
        "method": "dispatch",
        "params": {
            "topic": "secondTopic",
            "message": {
                "value": {
                    "body": "Message body"
                }
            }
        },
        "id": 2
    }
]
```

### From binaries

Build and run go application. Again, don't forget to set Apache Kafka brokers host by `MESSAGE_ROUTER_KAFKA_BROKERS` environment variable.

```bash
dep ensure
go build
./message-router
```

## Future tasks

* Acceptance testing
* Stats method
* Concurrency support
