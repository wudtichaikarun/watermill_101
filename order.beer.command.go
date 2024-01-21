package main

import (
	"context"
	"log"
	"math/rand"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/pkg/errors"
)

// OrderBeerHandler is a command handler, which handles OrderBeer command and emits BeerOrdered.
// BeerOrdered is not handled by any event handler, but we may use persistent Pub/Sub to handle it in the future.
type OrderBeerHandler struct {
	eventBus *cqrs.EventBus
}

func (o OrderBeerHandler) HandlerName() string {
	return "OrderBeerHandler"
}

func (o OrderBeerHandler) NewCommand() interface{} {
	return &OrderBeer{}
}

func (o OrderBeerHandler) Handle(ctx context.Context, c interface{}) error {
	cmd := c.(*OrderBeer)
	log.Printf("[receive] OrderBeerHandler receive command [orderBeerCmd] room: %s", cmd.RoomId)

	if rand.Int63n(10) == 0 {
		// sometimes there is no beer left, command will be retried
		return errors.Errorf("no beer left for room %s, please try later", cmd.RoomId)
	}

	log.Printf("[public] OrderBeerHandler public event [BeerOrdered] room: %s", cmd.RoomId)
	if err := o.eventBus.Publish(ctx, &BeerOrdered{
		RoomId: cmd.RoomId,
		Count:  cmd.Count,
	}); err != nil {
		return err
	}

	return nil
}
