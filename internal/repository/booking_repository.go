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

// Tạo một đặt phòng mới trong MongoDB
func CreateBooking(booking *model.Booking) error {
	collection := db.MongoClient.Database("testdb").Collection("bookings") // Thay thế "testdb" bằng tên database của bạn

	// Thiết lập thời gian đặt phòng cho booking
	booking.BookingDate = time.Now()

	// Thực hiện việc chèn một đặt phòng mới
	_, err := collection.InsertOne(context.TODO(), booking)
	if err != nil {
		return err
	}
	return nil
}

// Lấy tất cả đặt phòng của một người dùng từ MongoDB
func GetBookingsByUserID(userID primitive.ObjectID) ([]model.Booking, error) {
	var bookings []model.Booking
	collection := db.MongoClient.Database("testdb").Collection("bookings") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện query để lấy tất cả đặt phòng của người dùng
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var booking model.Booking
		if err := cursor.Decode(&booking); err != nil {
			log.Println("Decode error:", err)
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

// Cập nhật trạng thái một đặt phòng theo ID
func UpdateBooking(bookingID primitive.ObjectID, updatedBooking model.Booking) error {
	collection := db.MongoClient.Database("testdb").Collection("bookings") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc cập nhật đặt phòng
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": bookingID}, bson.M{"$set": updatedBooking})
	return err
}

// Xóa một đặt phòng theo ID
func DeleteBooking(bookingID primitive.ObjectID) error {
	collection := db.MongoClient.Database("testdb").Collection("bookings") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc xóa đặt phòng theo bookingID
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": bookingID})
	return err
}
