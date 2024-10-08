package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Favorite struct
type Favorite struct {
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`       // Foreign Key
	RoomID    primitive.ObjectID `bson:"room_id" json:"room_id"`       // Foreign Key
	CreatedAt time.Time          `bson:"created_at" json:"created_at"` // Time of favorite creation
}
