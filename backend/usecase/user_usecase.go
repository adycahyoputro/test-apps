package usecase

import (
	"backend/model"
	"backend/repository"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserUsecase interface {
	CreateUser(newUser *model.User) (*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) CreateUser(newUser *model.User) (*model.User, error) {
	// encripsi password
	//encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }
	if newUser.UserName == "" {
		return nil, errors.New("username is required")
	}
	if newUser.Email == "" {
		return nil, errors.New("email is required")
	}
	// newUser.Password = string(encryptedPassword)
	idUser, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to create uuid : %w", err)
	}
	newUser.ID = idUser.String()

	return u.userRepo.CreateUser(newUser)
}