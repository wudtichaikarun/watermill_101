package main

import (
	"context"
	"log"
	"math/rand"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

// OrderBeerOnRoomBooked is a event handler, which handles RoomBooked event and emits OrderBeer command.
type OrderBeerOnRoomBooked struct {
	commandBus *cqrs.CommandBus
}

func (o OrderBeerOnRoomBooked) HandlerName() string {
	// this name is passed to EventsSubscriberConstructor and used to generate queue name
	return "OrderBeerOnRoomBooked"
}

func (OrderBeerOnRoomBooked) NewEvent() interface{} {
	return &RoomBooked{}
}

func (o OrderBeerOnRoomBooked) Handle(ctx context.Context, e interface{}) error {
	event := e.(*RoomBooked)
	log.Printf("[receive] OrderBeerOnRoomBooked receive event [RoomBooked] room: %s", event.RoomId)

	orderBeerCmd := &OrderBeer{
		RoomId: event.RoomId,
		Count:  rand.Int63n(10) + 1,
	}

	log.Printf("[public] OrderBeerOnRoomBooked public command [orderBeerCmd] room: %s", event.RoomId)
	if err := o.commandBus.Send(ctx, orderBeerCmd); err != nil {
		return err
	}

	return nil

}
