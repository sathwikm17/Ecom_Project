package interfaces

import "github.com/sathwikm17/Ecom_Project/entities"

type IUser interface {
	Register(u *entities.User) (*entities.UserResponse, error)
	Login(u *entities.User) (string, error)
	LogOut() string
}
