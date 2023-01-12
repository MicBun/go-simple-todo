package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primary_key,auto_increment"`
	Title       string    `json:"title" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	Status      bool      `json:"status"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (t *Todo) CreateTodo(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) UpdateTodo(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var todo Todo
	err := db.Where("id = ? AND user_id = ?", t.ID, t.UserID).First(&todo).Error
	if err != nil {
		return err
	}
	err = db.Model(&todo).Updates(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) DeleteTodo(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var todo Todo
	err := db.Where("id = ?", t.ID).First(&todo).Error
	if err != nil {
		return err
	}
	err = db.Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) GetTodoByID(ctx *gin.Context) (Todo, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var todo Todo
	err := db.Where("id = ?", t.ID).First(&todo).Error
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (t *Todo) GetTodos(ctx *gin.Context) ([]Todo, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var todos []Todo
	err := db.Find(&todos).Error
	if err != nil {
		return []Todo{}, err
	}
	return todos, nil
}

func (t *Todo) GetTodosByUserID(ctx *gin.Context) ([]Todo, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var todos []Todo
	err := db.Where("user_id = ?", t.UserID).Find(&todos).Error
	if err != nil {
		return []Todo{}, err
	}
	return todos, nil
}

func (t *Todo) GetTodosByStatus(ctx *gin.Context) ([]Todo, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var todos []Todo
	err := db.Where("user_id = ? AND status = ?", t.UserID, t.Status).Find(&todos).Error
	if err != nil {
		return []Todo{}, err
	}
	return todos, nil
}
