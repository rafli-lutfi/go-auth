package service

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/go-auth/model"
	"github.com/rafli-lutfi/go-auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, user *model.UserLogin) (int, error)
	Register(ctx context.Context, user *model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserServcie(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func hashPassword(pwd *string) {
	bytePass := []byte(*pwd)
	hash, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)

	*pwd = string(hash)
}

func (u *userService) Login(ctx context.Context, user *model.UserLogin) (int, error) {
	dbUser, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return 0, err
	}

	if dbUser.Email == "" || dbUser.ID == 0 {
		return 0, errors.New("user not found")
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return 0, errors.New("password not matched")
	}

	return dbUser.ID, nil
}

func (u *userService) Register(ctx context.Context, user *model.User) (model.User, error) {
	dbUser, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return model.User{}, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return model.User{}, errors.New("email already exist")
	}

	hashPassword(&user.Password)

	dbUser, err = u.userRepository.CreateUser(ctx, *user)
	if err != nil {
		return model.User{}, err
	}

	return dbUser, nil
}
