package controllers

import (
	"github.com/MicBun/go-simple-todo/models"
	"github.com/MicBun/go-simple-todo/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateTodoInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// GetUserTodos godoc
// @Summary get user todos
// @Description Get all the todos related to the user
// @Tags todos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Security Bearer
// @Success 200 {array} models.Todo
// @Failure 404 {object} string
// @Router /todos [get]
func GetUserTodos(c *gin.Context) {
	id, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	user := models.User{ID: id}
	todos, err := user.GetTodos(c)
	if err != nil {
		c.JSON(404, gin.H{"error": "todos not found"})
		return
	}
	if len(todos) <= 0 {
		c.JSON(404, gin.H{"error": "no todo found"})
		return
	}
	c.JSON(200, gin.H{"todos": todos})
}

// CreateTodo godoc
// @Summary Create a new to-do
// @Description Create a new to-do and save it to the database
// @Tags todos
// @Accept  json
// @Produce  json
// @Param  todo body CreateTodoInput  true "Todo object"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /todos [post]
// @Security Bearer
func CreateTodo(c *gin.Context) {
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	todo := models.Todo{Title: input.Title, Description: input.Description, UserID: id}
	err = todo.CreateTodo(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "todo": todo})
}

// CompleteTodo godoc
// @Summary Complete a to-do
// @Description Complete a to-do by id
// @Tags todos
// @Accept  json
// @Produce  json
// @Param  id path int true "Todo ID"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /todos/{id} [put]
// @Security Bearer
func CompleteTodo(c *gin.Context) {
	userId, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(401, gin.H{"error": "wrong id format"})
		return
	}
	todo := models.Todo{UserID: userId, ID: uint(id), Status: true}
	err = todo.UpdateTodo(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "This todo is not yours"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "todo": todo})
}

// DeleteTodo godoc
// @Summary Delete a to-do
// @Description Delete a to-do by id
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /todos/{id} [delete]
// @Security Bearer
func DeleteTodo(c *gin.Context) {
	userId, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(401, gin.H{"error": "wrong id format"})
		return
	}
	todo := models.Todo{UserID: userId, ID: uint(id)}
	err = todo.DeleteTodo(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "This todo is not yours or not found"})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// GetTodoById godoc
// @Summary Get a to-do by id
// @Description Get a to-do by id
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} models.Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /todos/{id} [get]
// @Security Bearer
func GetTodoById(c *gin.Context) {
	userId, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(401, gin.H{"error": "wrong id format"})
		return
	}
	todoFind := models.Todo{UserID: userId, ID: uint(id)}
	todo, err := todoFind.GetTodoByID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "This todo is not yours or not found"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "todo": todo})
}

// GetUncompletedTodos godoc
// @Summary Get all uncompleted to-dos
// @Description Get all uncompleted to-dos
// @Tags todos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} []models.Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /todos/uncompleted [get]
// @Security Bearer
func GetUncompletedTodos(c *gin.Context) {
	userId, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	todoFind := models.Todo{UserID: userId, Status: false}
	todos, err := todoFind.GetTodosByStatus(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	if len(todos) == 0 {
		c.JSON(200, gin.H{"message": "success", "todos": "no todos found"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "todos": todos})
}

// GetCompletedTodos godoc
// @Summary Get all completed to-dos
// @Description Get all completed to-dos
// @Tags todos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} []models.Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /todos/completed [get]
// @Security Bearer
func GetCompletedTodos(c *gin.Context) {
	userId, err := jwtAuth.ExtractTokenID(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	todoFind := models.Todo{UserID: userId, Status: true}
	todos, err := todoFind.GetTodosByStatus(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	if len(todos) == 0 {
		c.JSON(200, gin.H{"message": "success", "todos": "no todos found"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "todos": todos})
}
