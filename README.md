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

## example log result

```
2024/01/21 16:19:42 [public] Guest public command [bookRoomCmd]
2024/01/21 16:19:42 [receive] BookRoomHandler receive command [bookRoomCmd]
2024/01/21 16:19:42 [public] BookRoomHandler public event [RoomBooked]
2024/01/21 16:19:42 [receive] OrderBeerOnRoomBooked receive event [RoomBooked]
2024/01/21 16:19:42 [receive] BookingsFinancialReport receive event [RoomBooked]
2024/01/21 16:19:42 generate report booked rooms for $230
2024/01/21 16:19:42 [public] OrderBeerOnRoomBooked public command [orderBeerCmd]
2024/01/21 16:19:42 [receive] OrderBeerHandler receive command [orderBeerCmd]
[watermill] 2024/01/21 16:19:42.330183 router.go:725:   level=ERROR msg="Handler returned error" err="no beer left for room 2, please try later"
2024/01/21 16:19:42 [receive] OrderBeerHandler receive command [orderBeerCmd]
2024/01/21 16:19:42 [public] OrderBeerHandler public event [BeerOrdered]
```
