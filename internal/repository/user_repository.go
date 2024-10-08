package repository

import (
	"context"
	"log"

	"begolang/db"
	"begolang/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Lấy tất cả người dùng từ MongoDB
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	collection := db.MongoClient.Database("testdb").Collection("users")

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

// Lấy người dùng theo ID
func GetUserByID(userID primitive.ObjectID) (*model.User, error) {
	var user model.User
	collection := db.MongoClient.Database("testdb").Collection("users")

	// Thực hiện query để tìm người dùng theo ID
	err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Trả về nil nếu không tìm thấy người dùng
		}
		return nil, err
	}

	return &user, nil
}

// Tạo một người dùng mới trong MongoDB
func CreateUser(user *model.User) error {
	collection := db.MongoClient.Database("testdb").Collection("users")

	// Thực hiện việc chèn một người dùng mới
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// Cập nhật một người dùng trong MongoDB
func UpdateUser(userID primitive.ObjectID, user *model.User) error {
	collection := db.MongoClient.Database("testdb").Collection("users")

	// Thực hiện việc cập nhật người dùng
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID}, // Tìm kiếm theo userID
		bson.M{"$set": user},  // Cập nhật các trường trong user
	)
	if err != nil {
		return err
	}
	return nil
}

// Xóa một người dùng trong MongoDB
func DeleteUser(userID primitive.ObjectID) error {
	collection := db.MongoClient.Database("testdb").Collection("users")

	// Thực hiện việc xóa người dùng
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": userID}) // Tìm kiếm theo userID
	if err != nil {
		return err
	}
	return nil
}
