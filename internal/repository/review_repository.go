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

// Tạo một đánh giá mới trong MongoDB
func CreateReview(review *model.Review) error {
	collection := db.MongoClient.Database("testdb").Collection("reviews") // Thay thế "testdb" bằng tên database của bạn

	// Thiết lập thời gian tạo cho đánh giá
	review.CreatedAt = time.Now()

	// Thực hiện việc chèn một đánh giá mới
	_, err := collection.InsertOne(context.TODO(), review)
	if err != nil {
		return err
	}
	return nil
}

// Lấy tất cả đánh giá của một phòng từ MongoDB
func GetReviewsByRoomID(roomID primitive.ObjectID) ([]model.Review, error) {
	var reviews []model.Review
	collection := db.MongoClient.Database("testdb").Collection("reviews") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện query để lấy tất cả đánh giá của phòng
	cursor, err := collection.Find(context.TODO(), bson.M{"room_id": roomID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var review model.Review
		if err := cursor.Decode(&review); err != nil {
			log.Println("Decode error:", err)
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}

// Cập nhật một đánh giá theo ID
func UpdateReview(reviewID primitive.ObjectID, updatedReview model.Review) error {
	collection := db.MongoClient.Database("testdb").Collection("reviews") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc cập nhật đánh giá
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": reviewID}, bson.M{"$set": updatedReview})
	return err
}

// Xóa một đánh giá theo ID
func DeleteReview(reviewID primitive.ObjectID) error {
	collection := db.MongoClient.Database("testdb").Collection("reviews") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc xóa đánh giá theo reviewID
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": reviewID})
	return err
}
