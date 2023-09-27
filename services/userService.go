package services

import (
	"github.com/sathwikm17/Ecom_Project/entities"
	"github.com/sathwikm17/Ecom_Project/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	User *mongo.Collection
}

func InitUserService(collection *mongo.Collection) interfaces.IUser {
	return &UserService{User: collection}
}

func (user *UserService) Register(u *entities.User) (*entities.UserResponse, error) {
	return nil, nil
}

func (user *UserService) Login(u *entities.User) (string, error) {
	return "", nil
}

func (user *UserService) LogOut() string {
	return ""
}
