package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`               // FK for the user
	RoomID       primitive.ObjectID `bson:"room_id" json:"room_id"`               // FK for the room
	BookingDate  time.Time          `bson:"booking_date" json:"booking_date"`     // Date of booking
	CheckInDate  time.Time          `bson:"check_in_date" json:"check_in_date"`   // Check-in date
	CheckOutDate time.Time          `bson:"check_out_date" json:"check_out_date"` // Check-out date
	Status       string             `bson:"status" json:"status"`                 // Status of the booking
}
