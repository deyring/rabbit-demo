# rabbit-demo
Demos for RabbitMQ use-cases.

## Prerequisites
- Golang
- Docker

## Run Rabbit Run
`docker compose up -d`

You can open the admin UI on `localhost:15672`.

User: `guest`

Password: `guest`

## Use Case 1: Worker Queue

Good fit for long running tasks, load balancing and scaling workloads.

To start the producer:

`go run ./worker-queue/producer/producer.go`

To start the workers:

`go run ./worker-queue/worker/worker.go`

Hint: to be able to add on the fly more workers to the queue, I added this line:

`ch.Qos(1, 0, true)`

Typically you would add some preFetch, to be more efficient on the network part.

## USe Case 2: Pub/Sub

Good fit for notifications / events.

To start the publisher:

`go run ./pub-sub/publisher/publisher.go`

To start the subscribers:

`go run ./pub-sub/sub1/subscriber.go`

`go run ./pub-sub/sub2/subscriber.go`