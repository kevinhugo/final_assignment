package repository

import (
	"assignment/app/models"
	"assignment/app/resource"
	"assignment/config"
	"errors"
)

type CommentRepository interface {
	AddComment(Comment *models.Comment, createData resource.InputComment) error
	GetComment(Comments *[]models.Comment, userId uint) error
	UpdateComment(Comment *models.Comment, createData resource.UpdateComment, userId uint) error
	DeleteComment(userId uint, id uint) error
}

func NewCommentRepository() CommentRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

// type Comment struct {
// 	gorm.Model
// 	CommentID      uint      `json:"order_id" gorm:"primary_key"`
// 	CustomerName string    `json:"customer_name"`
// 	CommentedAt    time.Time `json:"ordered_at" gorm:"autoCreateTime"`
// 	Items        []Item    `gorm:"foreignKey:CommentID"`
// }

// type Item struct {
// 	gorm.Model
// 	ItemID      uint   `json:"item_id" gorm:"primary_key"`
// 	ItemCode    string `json:"item_code" gorm:"index:document_user_id_index,unique"`
// 	Description string `json:"description" gorm:"index:document_user_id_index,unique"`
// 	Quantity    uint   `json:"quantity"`
// 	CommentID     uint   `json:"order_id"`
// 	Comment       Comment  `gorm:"foreignKey:CommentID"`
// }

func (db *dbConnection) AddComment(Comment *models.Comment, createData resource.InputComment) error {
	Comment.Message = createData.Message
	Comment.PhotoID = createData.PhotoID
	err := db.connection.Save(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) UpdateComment(Comment *models.Comment, createData resource.UpdateComment, userId uint) error {
	db.connection.Model(&Comment).Where("id = ?", Comment.ID).First(&Comment)
	if Comment.ID != 0 {
		if Comment.UserID != userId {
			return errors.New("Unauthorized to update this photo")
		}
	}
	Comment.Message = createData.Message
	err := db.connection.Save(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) GetComment(Comments *[]models.Comment, userId uint) error {
	db.connection.Model(Comments).Where("user_id = ?", userId).Preload("User").Preload("Photo").Find(Comments)
	return nil
}

// func (db *dbConnection) GetCommentList() ([]models.Comment, error, int64) {
// 	var Comment []models.Comment
// 	var count int64
// 	connection := db.connection.Model(&Comment).Preload("Items").Find(&Comment)
// 	err := connection.Error
// 	if err != nil {
// 		return Comment, err, 0
// 	}
// 	db.connection.Model(Comment).Count(&count)
// 	return Comment, nil, count
// }

// func (db *dbConnection) GetCommentDetailById(id uint, preload bool) (models.Comment, error) {
// 	var Comment models.Comment
// 	connection := db.connection
// 	fmt.Println("CommentId :", id)
// 	connection = connection.Where("id = ?", id)
// 	if preload {
// 		connection = connection.Preload("Items")
// 	}
// 	connection = connection.First(&Comment)
// 	err := connection.Error
// 	if err != nil {
// 		return Comment, err
// 	}
// 	return Comment, err
// }

func (db *dbConnection) DeleteComment(userId uint, id uint) error {
	var Comment models.Comment
	db.connection.Model(Comment).Where("id = ?", id).Find(&Comment)
	if Comment.ID != 0 {
		if Comment.UserID != userId {
			return errors.New("Unauthorized to delete this photo")
		}
	} else {
		return errors.New("Comment is gone")
	}
	err := db.connection.Unscoped().Delete(&Comment).Error
	if err != nil {
		return err
	}
	return nil
}
