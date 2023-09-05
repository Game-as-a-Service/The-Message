package models

import (
	"context"
	"fmt"

	"github.com/taimoor99/three-tier-golang/app/entities"
	"github.com/taimoor99/three-tier-golang/app/repositories/mongo"
	"github.com/taimoor99/three-tier-golang/utills"
)

type userRepo struct {
	Repo mongo.UserRepository
}

type UserModel interface {
	FindUserByID(ctx context.Context, userId string) (entities.Users, error)
	CreateUser(ctx context.Context, user entities.UsersCreateInput) (string, error)
	GetAllUsers(ctx context.Context, limit, offset int64) ([]entities.Users, error)
}

func NewUserModel(userRepository mongo.UserRepository) UserModel {
	return &userRepo{
		Repo: userRepository,
	}
}

func (m *userRepo) FindUserByID(ctx context.Context, userId string) (entities.Users, error) {
	user, err := m.Repo.FindUserById(ctx, userId)
	if err != nil {
		return entities.Users{}, err
	}
	return user, nil
}

func (m *userRepo) CreateUser(ctx context.Context, userReq entities.UsersCreateInput) (string, error) {
	user, err := m.Repo.FindUserByEmail(ctx, userReq.Email)
	if err == nil && user.Email != "" {
		return "", fmt.Errorf(utills.UserAlreadyExist)
	}

	user.Email = userReq.Email
	user.Name = userReq.Name
	user.Password = userReq.Password
	res, err := m.Repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return res.ID.Hex(), nil
}

func (m *userRepo) GetAllUsers(ctx context.Context, limit, offset int64) ([]entities.Users, error) {
	users, err := m.Repo.GetAllUsers(ctx, limit, offset)
	if err != nil {
		return []entities.Users{}, err
	}
	return users, nil
}
