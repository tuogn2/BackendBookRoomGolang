package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Payment struct
type Payment struct {
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`               // Foreign Key
	RoomID        primitive.ObjectID `bson:"room_id" json:"room_id"`               // Foreign Key
	BookingID     primitive.ObjectID `bson:"booking_id" json:"booking_id"`         // Foreign Key
	PaymentMethod string             `bson:"payment_method" json:"payment_method"` // Method of payment
	Amount        float64            `bson:"amount" json:"amount"`                 // Amount paid
	Status        string             `bson:"status" json:"status"`                 // Status of payment
}
