# rabbit-demo
Demos for RabbitMQ use-cases.

## Prerequisites
- Golang
- Docker

## Run Rabbit Run
There is a docker compose file in this repository. Essentially it brings up a rabbitMQ instance with default configuration, including the management UI.

You can simply start an instance in your local docker:
```bash
docker compose up -d
```

You can open the management UI on `localhost:15672`.

User: `guest`

Password: `guest`

## Use Case 1: Worker Queue

Good fit for long running tasks, load balancing and scaling workloads.

To start the producer:
```bash
go run ./worker-queue/producer/producer.go
```

To start the workers:
```bash
go run ./worker-queue/worker/worker.go
```

Hint: to be able to add on the fly more workers to the queue, I added this line:

`ch.Qos(1, 0, true)`

Typically you would add some preFetch, to be more efficient on the network part.

## Use Case 2: Pub/Sub

Good fit for notifications / events.

To start the publisher:
```bash
go run ./pub-sub/publisher/publisher.go
```

To start the subscribers:
```bash
go run ./pub-sub/sub1/subscriber.go
```
```bash
go run ./pub-sub/sub2/subscriber.go
```

Feel free to start multiple instances of sub1 or sub2. Checkout the behavior :-)

### Production [Non-]Suitability Disclaimer
Please keep in mind that this is just a demo / tutorial. The code misses proper tests and also does not cover some topics you should take into account in production environments. For example topics such as connection management, error handling, connection recovery, concurrency and metric collection are largely omitted for the sake of brevity. Such simplified code should not be considered production ready.