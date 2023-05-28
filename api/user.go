package api

import (
	"experimen_2/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(400, err)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(http.StatusOK, res)
		//fmt.Println("userLogin succ")
	} else {
		c.JSON(400, err)
		//fmt.Println("userLogin failed", err)
	}
}
