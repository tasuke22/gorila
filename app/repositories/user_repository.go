//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package repositories

import (
	"github.com/tasuke/go-mux-task/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	GetAllUserByEmail(user *[]models.User, email string) error
	CreateUser(user *models.User) error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

// Firstの中身の引数の話。interface{} は任意の型を受け入れることができる。可変長引数は0..nが入るから省略できる。
func (ur userRepository) GetUserByEmail(user *models.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur userRepository) GetAllUserByEmail(user *[]models.User, email string) error {
	if err := ur.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur userRepository) CreateUser(user *models.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
