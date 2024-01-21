package main

import (
	"context"
	fmt "fmt"
	"log"
	"sync"

	"github.com/theritikchoure/logx"
)

// BookingsFinancialReport is a read model, which calculates how much money we may earn from bookings.
// Like OrderBeerOnRoomBooked, it listens for RoomBooked event.
//
// This implementation is just writing to the memory. In production, you will probably will use some persistent storage.
type BookingsFinancialReport struct {
	handledBookings map[string]struct{}
	totalCharge     int64
	lock            sync.Mutex
}

func NewBookingsFinancialReport() *BookingsFinancialReport {
	return &BookingsFinancialReport{handledBookings: map[string]struct{}{}}
}

func (b BookingsFinancialReport) HandlerName() string {
	// this name is passed to EventsSubscriberConstructor and used to generate queue name
	return "BookingsFinancialReport"
}

func (BookingsFinancialReport) NewEvent() interface{} {
	return &RoomBooked{}
}

func (b *BookingsFinancialReport) Handle(ctx context.Context, e interface{}) error {

	// Handle may be called concurrently, so it need to be thread safe.
	b.lock.Lock()
	defer b.lock.Unlock()

	event := e.(*RoomBooked)
	eventMes := fmt.Sprintf("[receive] BookingsFinancialReport receive event [RoomBooked] room: %s", event.RoomId)
	logx.Log(eventMes, logx.FGWHITE, logx.BGGREEN)

	// When we are using Pub/Sub which doesn't provide exactly-once delivery semantics, we need to deduplicate messages.
	// GoChannel Pub/Sub provides exactly-once delivery,
	// but let's make this example ready for other Pub/Sub implementations.
	if _, ok := b.handledBookings[event.ReservationId]; ok {
		return nil
	}
	b.handledBookings[event.ReservationId] = struct{}{}

	b.totalCharge += event.Price

	log.Printf("[_______] BookingsFinancialReport generate report booked rooms: %s for $%d\n", event.RoomId, b.totalCharge)
	return nil
}
