package controllers

import (
	"github.com/MicBun/go-simple-todo/models"
	"github.com/MicBun/go-simple-todo/utils/jwtAuth"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary login to the system
// @Description Login to the system by providing the username and password
// @Tags authentication
// @Accept  json
// @Produce  json
// @Param  login body LoginInput true "Login object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userPassword := models.User{Username: input.Username, Password: input.Password}
	user, err := userPassword.LoginUser(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "username or password is incorrect."})
		return
	}
	jwtToken, err := jwtAuth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": jwtToken})
}

// Register godoc
// @Summary register to the system
// @Description Register to the system by providing the username and password
// @Tags authentication
// @Accept  json
// @Produce  json
// @Param   register     body    controllers.RegisterInput  true        "Register object"
// @Success 200 {object} models.User
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /register [post]
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newUser := models.User{Username: input.Username, Password: input.Password}
	err := newUser.CreateUser(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "username or password is incorrect."})
		return
	}
	c.JSON(200, gin.H{"message": "success", "user": newUser})
}
