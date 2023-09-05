package models

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/taimoor99/three-tier-golang/app/entities"
)

const email = "test@gmail.com"
const password = "test"

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetAllUsers(ctx context.Context, limit, offset int64) ([]entities.Users, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entities.Users), args.Error(1)
}

func (m *RepositoryMock) FindUserById(ctx context.Context, id string) (entities.Users, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Users), args.Error(1)
}

func (m *RepositoryMock) FindUserByEmail(ctx context.Context, email string) (entities.Users, error) {
	args := m.Called(email)
	return args.Get(0).(entities.Users), args.Error(1)
}

func (m *RepositoryMock) CreateUser(ctx context.Context, users entities.Users) (entities.Users, error) {
	args := m.Called(users)
	return args.Get(0).(entities.Users), args.Error(1)
}

var ums UserModel
var rm = new(RepositoryMock)

func init() {
	ums = NewUserModel(rm)
}

func TestUserModel_CreateUser(t *testing.T) {
	user := entities.Users{
		Name:     "test",
		Email:    email,
		Password: password,
	}
	userReq := entities.UsersCreateInput{
		Name:     "test",
		Email:    email,
		Password: password,
	}
	rm.On("FindUserByEmail", email).Return(entities.Users{}, nil)
	rm.On("CreateUser", user).Return(entities.Users{}, nil)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ums.CreateUser(ctx, userReq)
	assert.Nil(t, err)
}

func TestUserModel_FindUserByID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rm.On("FindUserById", "test").Return(entities.Users{Email: email}, nil)
	user, err := ums.FindUserByID(ctx, "test")
	assert.NoError(t, err)
	assert.Equal(t, user.Email, email)
}

func TestUserModel_GetAllUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rm.On("GetAllUsers", int64(10), int64(0)).Return([]entities.Users{}, nil)
	_, err := ums.GetAllUsers(ctx, 10, 0)
	assert.NoError(t, err)
}
