package events

import (
	"time"
)

// BookRoom represents a request to book a room
type BookRoom struct {
	RoomID    string    `json:"room_id"`
	GuestName string    `json:"guest_name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// RoomBooked represents a confirmation of a room booking
type RoomBooked struct {
	ReservationID string    `json:"reservation_id"`
	RoomID        string    `json:"room_id"`
	GuestName     string    `json:"guest_name"`
	Price         int64     `json:"price"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}

// OrderBeer represents a request to order beer
type OrderBeer struct {
	RoomID string `json:"room_id"`
	Count  int64  `json:"count"`
}

// BeerOrdered represents a confirmation of a beer order
type BeerOrdered struct {
	RoomID string `json:"room_id"`
	Count  int64  `json:"count"`
}
