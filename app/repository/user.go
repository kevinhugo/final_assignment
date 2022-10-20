package repository

import (
	"assignment/app/models"
	"assignment/app/resource"
	"assignment/config"
	"errors"

	"assignment/app/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	RegisterUser(User *models.User, createData resource.InputUser) error
	Login(email string, password string) (string, error)
	UpdateUser(User *models.User, createData resource.UpdateUser) error
	DeleteUser(id int) error
	HashPassword(text string) string
}

func NewUserRepository() UserRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

// type User struct {
// 	gorm.Model
// 	UserID      uint      `json:"order_id" gorm:"primary_key"`
// 	CustomerName string    `json:"customer_name"`
// 	UseredAt    time.Time `json:"ordered_at" gorm:"autoCreateTime"`
// 	Items        []Item    `gorm:"foreignKey:UserID"`
// }

// type Item struct {
// 	gorm.Model
// 	ItemID      uint   `json:"item_id" gorm:"primary_key"`
// 	ItemCode    string `json:"item_code" gorm:"index:document_user_id_index,unique"`
// 	Description string `json:"description" gorm:"index:document_user_id_index,unique"`
// 	Quantity    uint   `json:"quantity"`
// 	UserID     uint   `json:"order_id"`
// 	User       User  `gorm:"foreignKey:UserID"`
// }

func (db *dbConnection) RegisterUser(User *models.User, createData resource.InputUser) error {
	User.Username = createData.Username
	User.Email = createData.Email
	User.Age = createData.Age
	User.Password = db.HashPassword(createData.Password)
	err := db.connection.Save(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) UpdateUser(User *models.User, createData resource.UpdateUser) error {
	db.connection.Model(&User).Where("id = ?", User.ID).Find(&User)
	User.Username = createData.Username
	User.Email = createData.Email
	err := db.connection.Save(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) Login(email string, password string) (string, error) {
	var User models.User
	db.connection.Model(&User).Where("email = ?", email).Find(&User)
	if User.ID != 0 {
		if db.ComparePassword(password, User.Password) {
			token, err := utils.GenerateToken(User.ID)
			if err != nil {
				return "", errors.New("Something went wrong while trying to generate token.")
			}
			return token, nil
		} else {
			return "", errors.New("Email or password not match.")
		}
	}
	return "", errors.New("Email or password not match.")
}

// func (db *dbConnection) GetUserList() ([]models.User, error, int64) {
// 	var User []models.User
// 	var count int64
// 	connection := db.connection.Model(&User).Preload("Items").Find(&User)
// 	err := connection.Error
// 	if err != nil {
// 		return User, err, 0
// 	}
// 	db.connection.Model(User).Count(&count)
// 	return User, nil, count
// }

// func (db *dbConnection) GetUserDetailById(id uint, preload bool) (models.User, error) {
// 	var User models.User
// 	connection := db.connection
// 	fmt.Println("UserId :", id)
// 	connection = connection.Where("id = ?", id)
// 	if preload {
// 		connection = connection.Preload("Items")
// 	}
// 	connection = connection.First(&User)
// 	err := connection.Error
// 	if err != nil {
// 		return User, err
// 	}
// 	return User, err
// }

func (db *dbConnection) DeleteUser(id int) error {
	var User models.User
	err := db.connection.Unscoped().Delete(&User, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) HashPassword(text string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
	// hash := md5.Sum([]byte(text))
	// return hex.EncodeToString(hash[:])
}

func (db *dbConnection) ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == nil {
		return true
	}
	return false
}
