package httpserver

import (
	"Authentication/internaal/app"
	"Authentication/internaal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", health)
	userRepo := user.NewRepo(a.DB)
	userSvc := user.NewServices(userRepo, a.Config.JWTSecret)
	userhandler := user.NewHandler(userSvc)
	r.POST("/register", userhandler.Register)
	return r
}

