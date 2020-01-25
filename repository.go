package main

import (
	"github.com/jinzhu/gorm"
	auth "github.com/rickynyairo/plaeve-auth/proto/auth"
)

type Repository interface {
	GetAll() ([]*auth.User, error)
	Get(id string) (*auth.User, error)
	Create(user *auth.User) error
	GetByEmail(email string) (*auth.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*auth.User, error) {
	var users []*auth.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*auth.User, error) {
	var user *auth.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *auth.User) (*auth.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *auth.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetByEmail(email string) (*auth.User, error) {
	user := &auth.User{}
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
