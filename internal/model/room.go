package model

// Room đại diện cho một phòng trong hệ thống
type Room struct {
	Name        string   `bson:"name" json:"name"`               // Tên phòng
	Description string   `bson:"description" json:"description"` // Mô tả về phòng
	Price       float64  `bson:"price" json:"price"`             // Giá phòng
	Location    string   `bson:"location" json:"location"`       // Địa chỉ/phòng
	ListImage   []string `bson:"listimage" json:"listimage"`     // Danh sách đường dẫn tới hình ảnh của phòng
	Services    Services `bson:"services" json:"services"`       // Các dịch vụ kèm theo (bao gồm các tiện ích, dịch vụ và thông tin nhà vệ sinh)
	TypeRoom    string   `bson:"typeroom" json:"typeroom"`       // Loại phòng (ví dụ: Deluxe, Suite, ...)
	TypeErea    string   `bson:"typeerea" json:"typeerea"`       // Loại khu vực (ví dụ: thành phố, biển, núi,...)
	Quantity    int      `bson:"quantity" json:"quantity"`       // Số lượng phòng
}

// Services đại diện cho các dịch vụ và tiện ích của phòng
type Services struct {
	Facilities []string `bson:"facilities" json:"facilities"` // Danh sách các tiện ích (ví dụ: điều hòa, WiFi, ...)
	Service    []string `bson:"service" json:"service"`       // Các dịch vụ kèm theo (ví dụ: dọn phòng, giặt ủi,...)
	Bathroom   []string `bson:"bathroom" json:"bathroom"`     // Thông tin về nhà vệ sinh (ví dụ: có bồn tắm, có nước nóng,...)
}
