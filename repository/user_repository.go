package repository

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/go-auth/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetAllUser(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) GetUserByID(ctx context.Context, id int) (model.User, error) {
	user := model.User{}

	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, nil
	} else if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	user := model.User{}

	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, nil
	} else if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetAllUser(ctx context.Context) ([]model.User, error) {
	listOfUser := []model.User{}

	rows, err := u.db.WithContext(ctx).Find(&model.User{}).Rows()
	if err != nil {
		return []model.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		u.db.ScanRows(rows, &listOfUser)
	}

	return listOfUser, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user model.User) error {
	err := u.db.WithContext(ctx).Model(&model.User{}).Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	err := u.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}

	return nil
}
