package controller

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"

	"github.com/mopeneko/novel-gamest/api/domain"
	"github.com/mopeneko/novel-gamest/api/interface/database"
)

// UserController is an interface for User Use Case
type UserController struct {
	UserRepository database.UserRepository
}

// Create new user
func (controller *UserController) Create(id, password string) (domain.User, error) {
	// パスワードをハッシュ化
	hashedPassword := sha512.Sum512([]byte(password))
	hashedHexPassword := hex.EncodeToString(hashedPassword[:])

	user := domain.User{
		UserID:   id,
		Password: hashedHexPassword,
	}
	err := controller.UserRepository.Save(user)
	if err != nil {
		return user, err
	}

	return controller.GetByID(id)
}

// GetByID returns an user which has same ID
func (controller *UserController) GetByID(id string) (domain.User, error) {
	return controller.UserRepository.FindByID(id)
}

// GetWithAuthentication returns an user which has same ID and password
func (controller *UserController) GetWithAuthentication(id, password string) (domain.User, error) {
	user, err := controller.GetByID(id)
	if err != nil {
		return user, err
	}

	// パスワードをハッシュ化
	hashedPassword := sha512.Sum512([]byte(password))
	hashedHexPassword := hex.EncodeToString(hashedPassword[:])

	if user.Password != hashedHexPassword {
		return user, errors.New("パスワードが間違っています。")
	}

	return user, nil
}
