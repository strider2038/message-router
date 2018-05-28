# Request validation

## Rules for valid request

* url = "/"
* http method = POST
* content type = application/json
* body != nil

## POST request body

[
    {
        topic: string
        message: {
            headers: {
                message_id: string
                timestamp: float (unix timestamp)
                serializationId: string
            }
            properties: { ... }
            body: { ... }
        }
    }
]
