package router

import (
	"experimen_2/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	//cookie
	store := cookie.NewStore([]byte("something-very-secret"))
	engine.Use(sessions.Sessions("mysession", store))
	v1 := engine.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		//报错 因为Api中的方法和这里的不一致
		v1.POST("user/login", api.UserLogin)
	}

	return engine
}
