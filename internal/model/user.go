package model

import (
	"time"
)

type User struct {
	Name      string    `json:"name"`                // Tên người dùng
	Email     string    `json:"email" gorm:"unique"` // Email duy nhất
	Password  string    `json:"password"`            // Mật khẩu người dùng
	Phone     string    `json:"phone"`               // Số điện thoại người dùng
	Avatar    string    `json:"avt"`                 // Đường dẫn tới ảnh đại diện
	Birthday  time.Time `json:"birthday"`            // Ngày sinh
	CreatedAt time.Time `json:"created_at"`          // Thời gian tạo
	UpdatedAt time.Time `json:"updated_at"`          // Thời gian cập nhật
}
