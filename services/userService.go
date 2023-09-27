package services

import (
	"context"
	"strings"
	"time"

	"github.com/sathwikm17/Ecom_Project/entities"
	"github.com/sathwikm17/Ecom_Project/interfaces"
	"github.com/sathwikm17/Ecom_Project/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	User *mongo.Collection
}

func InitUserService(collection *mongo.Collection) interfaces.IUser {
	return &UserService{User: collection}
}

func (user *UserService) Register(u *entities.User) (*entities.UserResponse, error) {
	ctx := context.Background()
	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt
	u.Email = strings.ToLower(u.Email)

	hashedPassword, _ := utils.HashPassword(u.Password)
	u.Password = hashedPassword

	_, err := user.User.InsertOne(ctx, &u)
	if err != nil {
		return nil, err
	}

	registeredUser := entities.UserResponse{Response: "User registered successfully"}

	return &registeredUser, nil
}

func (user *UserService) Login(u *entities.User) (string, error) {
	return "", nil
}

func (user *UserService) LogOut() string {
	return ""
}
