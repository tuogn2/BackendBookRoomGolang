package repository

import (
	"context"
	"log"

	"begolang/db"
	"begolang/internal/model"

	"go.mongodb.org/mongo-driver/bson"
)

// Lấy tất cả phòng từ MongoDB
func GetAllRooms() ([]model.Room, error) {
	var rooms []model.Room
	collection := db.MongoClient.Database("testdb").Collection("rooms") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện query để lấy tất cả phòng
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var room model.Room
		if err := cursor.Decode(&room); err != nil {
			log.Println("Decode error:", err)
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

// Tạo một phòng mới trong MongoDB
func CreateRoom(room *model.Room) error {
	collection := db.MongoClient.Database("testdb").Collection("rooms") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc chèn một phòng mới
	_, err := collection.InsertOne(context.TODO(), room)
	if err != nil {
		return err
	}
	return nil
}

// Lấy phòng theo ID từ MongoDB
func GetRoomByID(id string) (*model.Room, error) {
	var room model.Room
	collection := db.MongoClient.Database("testdb").Collection("rooms") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc tìm kiếm phòng theo ID
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&room)
	if err != nil {
		return nil, err // Trả về nil nếu không tìm thấy phòng
	}

	return &room, nil // Trả về phòng tìm thấy
}

// Cập nhật thông tin phòng trong MongoDB
func UpdateRoom(id string, room *model.Room) error {
	collection := db.MongoClient.Database("testdb").Collection("rooms") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc cập nhật phòng theo ID
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": room})
	return err
}

// Xóa phòng theo ID từ MongoDB
func DeleteRoom(id string) error {
	collection := db.MongoClient.Database("testdb").Collection("rooms") // Thay thế "testdb" bằng tên database của bạn

	// Thực hiện việc xóa phòng theo ID
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
