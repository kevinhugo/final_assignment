package repository

import (
	"assignment/app/models"
	"assignment/app/resource"
	"assignment/config"
	"errors"
)

type SocialMediaRepository interface {
	AddSocialMedia(SocialMedia *models.SocialMedia, createData resource.InputSocialMedia) error
	GetSocialMedia(SocialMedias *[]models.SocialMedia, userId uint) error
	UpdateSocialMedia(SocialMedia *models.SocialMedia, createData resource.UpdateSocialMedia, userId uint) error
	DeleteSocialMedia(userId uint, id uint) error
}

func NewSocialMediaRepository() SocialMediaRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

// type SocialMedia struct {
// 	gorm.Model
// 	SocialMediaID      uint      `json:"order_id" gorm:"primary_key"`
// 	CustomerName string    `json:"customer_name"`
// 	SocialMediaedAt    time.Time `json:"ordered_at" gorm:"autoCreateTime"`
// 	Items        []Item    `gorm:"foreignKey:SocialMediaID"`
// }

// type Item struct {
// 	gorm.Model
// 	ItemID      uint   `json:"item_id" gorm:"primary_key"`
// 	ItemCode    string `json:"item_code" gorm:"index:document_user_id_index,unique"`
// 	Description string `json:"description" gorm:"index:document_user_id_index,unique"`
// 	Quantity    uint   `json:"quantity"`
// 	SocialMediaID     uint   `json:"order_id"`
// 	SocialMedia       SocialMedia  `gorm:"foreignKey:SocialMediaID"`
// }

func (db *dbConnection) AddSocialMedia(SocialMedia *models.SocialMedia, createData resource.InputSocialMedia) error {
	SocialMedia.Name = createData.Name
	SocialMedia.SocialMedialUrl = createData.SocialMedialUrl
	err := db.connection.Save(SocialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) UpdateSocialMedia(SocialMedia *models.SocialMedia, createData resource.UpdateSocialMedia, userId uint) error {
	db.connection.Model(&SocialMedia).Where("id = ?", SocialMedia.ID).First(&SocialMedia)
	if SocialMedia.ID != 0 {
		if SocialMedia.UserID != userId {
			return errors.New("Unauthorized to update this social media")
		}
	}
	SocialMedia.Name = createData.Name
	SocialMedia.SocialMedialUrl = createData.SocialMedialUrl
	err := db.connection.Save(SocialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) GetSocialMedia(SocialMedias *[]models.SocialMedia, userId uint) error {
	db.connection.Model(SocialMedias).Where("user_id = ?", userId).Preload("User").Preload("Photo").Find(SocialMedias)
	return nil
}

// func (db *dbConnection) GetSocialMediaList() ([]models.SocialMedia, error, int64) {
// 	var SocialMedia []models.SocialMedia
// 	var count int64
// 	connection := db.connection.Model(&SocialMedia).Preload("Items").Find(&SocialMedia)
// 	err := connection.Error
// 	if err != nil {
// 		return SocialMedia, err, 0
// 	}
// 	db.connection.Model(SocialMedia).Count(&count)
// 	return SocialMedia, nil, count
// }

// func (db *dbConnection) GetSocialMediaDetailById(id uint, preload bool) (models.SocialMedia, error) {
// 	var SocialMedia models.SocialMedia
// 	connection := db.connection
// 	fmt.Println("SocialMediaId :", id)
// 	connection = connection.Where("id = ?", id)
// 	if preload {
// 		connection = connection.Preload("Items")
// 	}
// 	connection = connection.First(&SocialMedia)
// 	err := connection.Error
// 	if err != nil {
// 		return SocialMedia, err
// 	}
// 	return SocialMedia, err
// }

func (db *dbConnection) DeleteSocialMedia(userId uint, id uint) error {
	var SocialMedia models.SocialMedia
	db.connection.Model(SocialMedia).Where("id = ?", id).Find(&SocialMedia)
	if SocialMedia.ID != 0 {
		if SocialMedia.UserID != userId {
			return errors.New("Unauthorized to delete this social media")
		}
	} else {
		return errors.New("SocialMedia is gone")
	}
	err := db.connection.Unscoped().Delete(&SocialMedia).Error
	if err != nil {
		return err
	}
	return nil
}
