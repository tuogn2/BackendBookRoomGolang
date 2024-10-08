package repository

import (
	"context"
	"log"

	"begolang/db"
	"begolang/internal/model"

	"go.mongodb.org/mongo-driver/bson"
)

// Lấy tất cả người dùng từ MongoDB
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	collection := db.MongoClient.Database("testdb").Collection("users") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện query để lấy tất cả người dùng
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Decode error:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Tạo một người dùng mới trong MongoDB
func CreateUser(user *model.User) error {
	collection := db.MongoClient.Database("testdb").Collection("users") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc chèn một người dùng mới
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}
