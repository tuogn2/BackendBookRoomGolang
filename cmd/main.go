package main

import (
	"begolang/db"
	"begolang/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Kết nối với database
	db.InitMongoDB()

	// Khởi tạo router
	r := mux.NewRouter()

	// Định nghĩa các route
	r.HandleFunc("/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")

	// Khởi chạy server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
