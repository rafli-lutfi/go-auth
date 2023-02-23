package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rafli-lutfi/go-auth/model"
	"github.com/rafli-lutfi/go-auth/service"
)

type UserAPI interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func generateJWT(id int) string {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenString
}

func (u *userAPI) Login(c *gin.Context) {
	var creds model.UserLogin

	err := c.ShouldBindJSON(&creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	if creds.Email == "" || creds.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "user or password is empty",
			"data":    "",
		})
		return
	}

	userID, err := u.userService.Login(c.Request.Context(), &creds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	tokenString := generateJWT(userID)

	c.SetCookie("token", tokenString, 3600*24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "login success",
		"data": gin.H{
			"id": userID,
		},
	})
}

func (u *userAPI) Register(c *gin.Context) {
	var creds model.UserRegister

	err := c.ShouldBindJSON(&creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	if creds.Email == "" || creds.Password == "" || creds.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "register data is empty",
			"data":    "",
		})
		return
	}

	user := model.User{
		Username: creds.Username,
		Email:    creds.Email,
		Password: creds.Password,
	}

	newUser, err := u.userService.Register(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "success create new user",
		"data": gin.H{
			"username": newUser.Username,
		},
	})
}

func (u *userAPI) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "logout success",
		"data":    "",
	})
}
