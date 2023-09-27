package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathwikm17/Ecom_Project/entities"
	"github.com/sathwikm17/Ecom_Project/interfaces"
)

type UserController struct {
	UserService interfaces.IUser
}

func InitUserController(usrService interfaces.IUser) *UserController {
	return &UserController{UserService: usrService}
}

func (user *UserController) Register(c *gin.Context) {
	var usr entities.User

	err := c.BindJSON(&usr)
	if err != nil {
		return
	}

	result, err := user.UserService.Register(&usr)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
}
