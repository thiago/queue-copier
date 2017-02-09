Queue Copier
==========

Move messages from one SQS queue to another.

## Installation
```bash
$ go get github.com/DispatchMe/queue-copier
```

## Usage MODES

### SQS
```bash
$ ./queue-copier
-mode="sqs"
-queue1="<url for origin>" \
-queue2="<url for destination>"
```

### RabbitMQ
```bash
$ RABBITMQ_CONNECTION_STRING="<url connection string>" ./queue-copier
-mode="rabbit"
-dead-queue="name-of-dead-queue" \
-exchange="name-of-exchange"
```
