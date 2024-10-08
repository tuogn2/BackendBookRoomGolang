package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Review struct
type Review struct {
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`       // Foreign Key
	RoomID    primitive.ObjectID `bson:"room_id" json:"room_id"`       // Foreign Key
	Rating    int                `bson:"rating" json:"rating"`         // Rating given by user
	Comment   string             `bson:"comment" json:"comment"`       // Review comment
	CreatedAt time.Time          `bson:"created_at" json:"created_at"` // Time of review creation
}
