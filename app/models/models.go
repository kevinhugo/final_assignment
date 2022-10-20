package models

import (
	"time"

	"gorm.io/gorm"
)

// type Document struct {
// 	gorm.Model
// 	ID               uint   `json:"id" gorm:"primary_key"`
// 	NotificationID       uint   `json:"NotificationID"  binding:"required" gorm:"index:,unique,composite:orderNotification"`
// 	UserID          string `json:"UserID" binding:"required" gorm:"index:,unique,composite:orderNotification"`
// 	Notification         Notification
// 	RefferenceNumber string             `json:"refferenceNumber"`
// 	TotalAmount      int                `json:"totalAmount" binding:"required"`
// 	Items            []DocumentUser `gorm:"foreignKey:DocumentID"`
// 	Status           string             `json:"status" gorm:"default:pending"`
// 	SelectedPayment  string             `json:"selectedPayment"`
// 	Token            string             `json:"token"`
// 	TokenAuth        string             `json:"tokenAuth"`
// }

// type DocumentUser struct {
// 	gorm.Model
// 	DocumentID uint
// 	Name          string `json:"name"`
// 	Price         int    `json:"price"`
// }

type User struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}

type Photo struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}

type Comment struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	PhotoID   uint      `json:"photo_id"`
	Photo     Photo     `gorm:"foreignKey:PhotoID"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}

type SocialMedia struct {
	gorm.Model
	ID              uint      `json:"id" gorm:"primary_key"`
	Name            string    `json:"name"`
	SocialMedialUrl string    `json:"social_media_url"`
	UserID          uint      `json:"user_id"`
	User            User      `gorm:"foreignKey:UserID"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}
