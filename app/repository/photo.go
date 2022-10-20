package repository

import (
	"assignment/app/models"
	"assignment/app/resource"
	"assignment/config"
	"errors"
)

type PhotoRepository interface {
	AddPhoto(Photo *models.Photo, createData resource.InputPhoto) error
	GetPhoto(Photos *[]models.Photo, userId uint) error
	UpdatePhoto(Photo *models.Photo, createData resource.InputPhoto, userId uint) error
	DeletePhoto(userId uint, id uint) error
}

func NewPhotoRepository() PhotoRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

// type Photo struct {
// 	gorm.Model
// 	PhotoID      uint      `json:"order_id" gorm:"primary_key"`
// 	CustomerName string    `json:"customer_name"`
// 	PhotoedAt    time.Time `json:"ordered_at" gorm:"autoCreateTime"`
// 	Items        []Item    `gorm:"foreignKey:PhotoID"`
// }

// type Item struct {
// 	gorm.Model
// 	ItemID      uint   `json:"item_id" gorm:"primary_key"`
// 	ItemCode    string `json:"item_code" gorm:"index:document_user_id_index,unique"`
// 	Description string `json:"description" gorm:"index:document_user_id_index,unique"`
// 	Quantity    uint   `json:"quantity"`
// 	PhotoID     uint   `json:"order_id"`
// 	Photo       Photo  `gorm:"foreignKey:PhotoID"`
// }

func (db *dbConnection) AddPhoto(Photo *models.Photo, createData resource.InputPhoto) error {
	Photo.Title = createData.Title
	Photo.Caption = createData.Caption
	Photo.PhotoUrl = createData.PhotoUrl
	err := db.connection.Save(Photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) UpdatePhoto(Photo *models.Photo, createData resource.InputPhoto, userId uint) error {
	db.connection.Model(&Photo).Where("id = ?", Photo.ID).First(&Photo)
	if Photo.ID != 0 {
		if Photo.UserID != userId {
			return errors.New("Unauthorized to update this photo")
		}
	}
	Photo.Title = createData.Title
	Photo.Caption = createData.Caption
	Photo.PhotoUrl = createData.PhotoUrl
	err := db.connection.Save(Photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) GetPhoto(Photos *[]models.Photo, userId uint) error {
	db.connection.Model(Photos).Where("user_id = ?", userId).Find(Photos)
	return nil
}

// func (db *dbConnection) GetPhotoList() ([]models.Photo, error, int64) {
// 	var Photo []models.Photo
// 	var count int64
// 	connection := db.connection.Model(&Photo).Preload("Items").Find(&Photo)
// 	err := connection.Error
// 	if err != nil {
// 		return Photo, err, 0
// 	}
// 	db.connection.Model(Photo).Count(&count)
// 	return Photo, nil, count
// }

// func (db *dbConnection) GetPhotoDetailById(id uint, preload bool) (models.Photo, error) {
// 	var Photo models.Photo
// 	connection := db.connection
// 	fmt.Println("PhotoId :", id)
// 	connection = connection.Where("id = ?", id)
// 	if preload {
// 		connection = connection.Preload("Items")
// 	}
// 	connection = connection.First(&Photo)
// 	err := connection.Error
// 	if err != nil {
// 		return Photo, err
// 	}
// 	return Photo, err
// }

func (db *dbConnection) DeletePhoto(userId uint, id uint) error {
	var Photo models.Photo
	db.connection.Model(Photo).Where("id = ?", id).Find(&Photo)
	if Photo.ID != 0 {
		if Photo.UserID != userId {
			return errors.New("Unauthorized to delete this photo")
		}
	} else {
		return errors.New("Photo is gone")
	}
	err := db.connection.Unscoped().Delete(&Photo).Error
	if err != nil {
		return err
	}
	return nil
}
