package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	Todo      []Todo    `json:"todo" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *User) LoginUser(ctx *gin.Context) (User, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var user User
	err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) CreateUser(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	err := db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var user User
	err := db.Where("id = ?", u.ID).First(&user).Error
	if err != nil {
		return err
	}
	err = db.Model(&user).Updates(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteUser(ctx *gin.Context) error {
	db := ctx.MustGet("db").(*gorm.DB)
	var user User
	err := db.Where("id = ?", u.ID).First(&user).Error
	if err != nil {
		return err
	}
	err = db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetTodos(ctx *gin.Context) ([]Todo, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var todos []Todo
	err := db.Where("user_id = ?", u.ID).Find(&todos).Error
	if err != nil {
		return []Todo{}, err
	}
	return todos, nil
}

func (u *User) GetTodoByID(ctx *gin.Context, id uint) (Todo, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var todo Todo
	err := db.Where("user_id = ? AND id = ?", u.ID, id).First(&todo).Error
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}
