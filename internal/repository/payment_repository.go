package repository

import (
	"context"
	"log"

	"begolang/db"
	"begolang/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tạo một thanh toán mới trong MongoDB
func CreatePayment(payment *model.Payment) error {
	collection := db.MongoClient.Database("testdb").Collection("payments") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc chèn một thanh toán mới
	_, err := collection.InsertOne(context.TODO(), payment)
	if err != nil {
		return err
	}
	return nil
}

// Lấy tất cả thanh toán của một người dùng từ MongoDB
func GetPaymentsByUserID(userID primitive.ObjectID) ([]model.Payment, error) {
	var payments []model.Payment
	collection := db.MongoClient.Database("testdb").Collection("payments") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện query để lấy tất cả thanh toán của người dùng
	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var payment model.Payment
		if err := cursor.Decode(&payment); err != nil {
			log.Println("Decode error:", err)
			return nil, err
		}
		payments = append(payments, payment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}

// Cập nhật trạng thái một thanh toán theo ID
func UpdatePayment(paymentID primitive.ObjectID, updatedPayment model.Payment) error {
	collection := db.MongoClient.Database("testdb").Collection("payments") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc cập nhật thanh toán
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": paymentID}, bson.M{"$set": updatedPayment})
	return err
}

// Xóa một thanh toán theo ID
func DeletePayment(paymentID primitive.ObjectID) error {
	collection := db.MongoClient.Database("testdb").Collection("payments") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc xóa thanh toán theo paymentID
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": paymentID})
	return err
}
