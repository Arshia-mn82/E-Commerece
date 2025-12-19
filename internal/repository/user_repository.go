package repository

import (
	"E-Commerce/internal/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(u *model.User) error {
	return ur.db.Create(u).Error
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	var u model.User

	err := ur.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (ur *UserRepository) FindByID(ID uint) (*model.User, error) {
	var u model.User
	err := ur.db.First(&u, ID).Error

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
