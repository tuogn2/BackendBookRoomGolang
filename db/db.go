package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongoDB() {
	// MongoDB URI (đảm bảo thay thế bằng URI của bạn, có thể từ MongoDB Atlas hoặc cài đặt cục bộ)
	uri := "mongodb+srv://tuong:021603@golang.f1ebz.mongodb.net/?retryWrites=true&w=majority&appName=golang"

	// Thiết lập các tùy chọn kết nối
	clientOptions := options.Client().ApplyURI(uri)

	// Tạo context với timeout để kết nối MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Kết nối tới MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Kiểm tra kết nối
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	MongoClient = client
	fmt.Println("Connected to MongoDB!")
}
