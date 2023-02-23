package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafli-lutfi/go-auth/controller/api"
	"github.com/rafli-lutfi/go-auth/middleware"
	"github.com/rafli-lutfi/go-auth/repository"
	"github.com/rafli-lutfi/go-auth/service"
	"gorm.io/gorm"
)

type APIHandler struct {
	UserAPI api.UserAPI
}

func RunServer(db *gorm.DB, r *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserServcie(userRepository)
	userAPI := api.NewUserAPI(userService)

	handler := APIHandler{
		UserAPI: userAPI,
	}

	server := r.Group("/api/v1")

	public := server.Group("/public")
	user := public.Group("/user")
	user.POST("/login", handler.UserAPI.Login)
	user.POST("/register", handler.UserAPI.Register)
	user.GET("/logout", middleware.Authorization, handler.UserAPI.Logout)
}
