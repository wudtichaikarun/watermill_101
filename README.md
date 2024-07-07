# Example Golang CQRS application

This application is using [Watermill CQRS](http://watermill.io/docs/cqrs) component.

Detailed documentation for CQRS can be found in Watermill's docs: [http://watermill.io/docs/cqrs#usage](http://watermill.io/docs/cqrs).

### Usage

Example domain
As an example, we will use a simple domain, that is responsible for handing room booking in a hotel.

We will use Event Storming notation to show the model of this domain.

Legend:

- blue post-its are commands
- orange post-its are events
- green post-its are read models, asynchronously generated from events
- violet post-its are policies, which are triggered by events and produce commands
- pink post its are hot-spots; we mark places where problems often occur

![CQRS Event Storming](https://threedots.tech/watermill-io/cqrs-example-storming.png)

## install protobuf

```
$ brew install protobuf

// check version
$ protoc --version
```

## generate protobuf

```
$ protoc --go_out=. inputs/events.proto

```

## rabbitmq

local
https://www.rabbitmq.com/docs/download

```
// username: guest password: guest
// main.go => var amqpAddress = "amqp://guest:guest@localhost:5672/"
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management
```

using docker compose

```
// main.og -> var amqpAddress = "amqp://guest:guest@rabbitmq:5672/"
make compose-dev
```

## example log result

```
2024/01/21 16:53:22 [public] Guest public command [bookRoomCmd] room: 2
2024/01/21 16:53:22 [receive] BookRoomHandler receive command [bookRoomCmd] room: 2
2024/01/21 16:53:22 [public] BookRoomHandler public event [RoomBooked] room: 2
2024/01/21 16:53:22 [receive] BookingsFinancialReport receive event [RoomBooked] room: 2
2024/01/21 16:53:22 [_______] BookingsFinancialReport generate report booked rooms: 2 for $230
2024/01/21 16:53:22 [receive] OrderBeerOnRoomBooked receive event [RoomBooked] room: 2
2024/01/21 16:53:22 [public] OrderBeerOnRoomBooked public command [orderBeerCmd] room: 2
2024/01/21 16:53:22 [receive] OrderBeerHandler receive command [orderBeerCmd] room: 2
[watermill] 2024/01/21 16:53:22.920800 router.go:725:   level=ERROR msg="Handler returned error" err="no beer left for room 2, please try later"
2024/01/21 16:53:22 [receive] OrderBeerHandler receive command [orderBeerCmd] room: 2
2024/01/21 16:53:22 [public] OrderBeerHandler public event [BeerOrdered] room: 2
```

- blue for command
- green for event

![result image](./images/result.png)
