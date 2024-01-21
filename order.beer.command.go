package main

import (
	"context"
	fmt "fmt"
	"math/rand"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/pkg/errors"
	"github.com/theritikchoure/logx"
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

	cmdMes := fmt.Sprintf("[receive] OrderBeerHandler receive command [orderBeerCmd] room: %s", cmd.RoomId)
	logx.Log(cmdMes, logx.FGWHITE, logx.BGBLUE)

	if rand.Int63n(10) == 0 {
		// sometimes there is no beer left, command will be retried
		return errors.Errorf("no beer left for room %s, please try later", cmd.RoomId)
	}

	eventMes := fmt.Sprintf("[public] OrderBeerHandler public event [BeerOrdered] room: %s", cmd.RoomId)
	logx.Log(eventMes, logx.FGWHITE, logx.BGGREEN)

	if err := o.eventBus.Publish(ctx, &BeerOrdered{
		RoomId: cmd.RoomId,
		Count:  cmd.Count,
	}); err != nil {
		return err
	}

	return nil
}
