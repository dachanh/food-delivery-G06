package usermodel

import (
	"github.com/dachanh/food-delivery-G06/common"
)

type User struct {
	common.SQLModel
	Email     string        `json:"email"gorm:"column:email;"`
	Password  string        `json:"-" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"last_name"gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"role"gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u User) TableName() string { return "users" }

type UserCreate struct {
	common.SQLModel
	Email     string        `json:"email"gorm:"column:email;"`
	Password  string        `json:"-" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"last_name"gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"role"gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u UserCreate) TableName() string { return User{}.TableName() }
