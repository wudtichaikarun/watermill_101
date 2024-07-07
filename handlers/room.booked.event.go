package handlers

import (
	"context"
	fmt "fmt"
	"math/rand"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/theritikchoure/logx"
	"github.com/wudtichaikarun/watermill_101/pkg/events"
)

// OrderBeerOnRoomBooked is a event handler, which handles RoomBooked event and emits OrderBeer command.
type OrderBeerOnRoomBooked struct {
	CommandBus *cqrs.CommandBus
}

func (o OrderBeerOnRoomBooked) HandlerName() string {
	// this name is passed to EventsSubscriberConstructor and used to generate queue name
	return "OrderBeerOnRoomBooked"
}

func (OrderBeerOnRoomBooked) NewEvent() interface{} {
	return &events.RoomBooked{}
}

func (o OrderBeerOnRoomBooked) Handle(ctx context.Context, e interface{}) error {
	event := e.(*events.RoomBooked)

	eventMes := fmt.Sprintf("[receive] OrderBeerOnRoomBooked receive event [RoomBooked] room: %s", event.RoomID)
	logx.Log(eventMes, logx.FGWHITE, logx.BGGREEN)

	orderBeerCmd := &events.OrderBeer{
		RoomID: event.RoomID,
		Count:  rand.Int63n(10) + 1,
	}

	cmdMes := fmt.Sprintf("[public] OrderBeerOnRoomBooked public command [orderBeerCmd] room: %s", event.RoomID)
	logx.Log(cmdMes, logx.FGWHITE, logx.BGBLUE)

	if err := o.CommandBus.Send(ctx, orderBeerCmd); err != nil {
		return err
	}

	return nil

}
