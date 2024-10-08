package repository

import (
	"context"
	"log"
	"time"

	"begolang/db"
	"begolang/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tạo một yêu thích mới trong MongoDB
func CreateFavorite(favorite *model.Favorite) error {
	collection := db.MongoClient.Database("testdb").Collection("favorites") // Thay thế "testdb" bằng tên database của bạn

	// Thiết lập thời gian tạo cho yêu thích
	favorite.CreatedAt = time.Now()

	// Thực hiện việc chèn một yêu thích mới
	_, err := collection.InsertOne(context.TODO(), favorite)
	if err != nil {
		return err
	}
	return nil
}

// Lấy tất cả yêu thích của một người dùng từ MongoDB
func GetFavoritesByUserID(userID primitive.ObjectID) ([]model.Favorite, error) {
	var favorites []model.Favorite
	collection := db.MongoClient.Database("testdb").Collection("favorites") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện query để lấy tất cả yêu thích của người dùng
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var favorite model.Favorite
		if err := cursor.Decode(&favorite); err != nil {
			log.Println("Decode error:", err)
			return nil, err
		}
		favorites = append(favorites, favorite)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return favorites, nil
}

// Xóa một yêu thích theo ID người dùng và ID phòng
func DeleteFavorite(userID, roomID primitive.ObjectID) error {
	collection := db.MongoClient.Database("testdb").Collection("favorites") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc xóa yêu thích theo userID và roomID
	_, err := collection.DeleteOne(context.TODO(), bson.M{
		"user_id": userID,
		"room_id": roomID,
	})
	return err
}
