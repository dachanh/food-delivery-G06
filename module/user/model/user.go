package usermodel

import (
	"errors"
	"github.com/dachanh/food-delivery-G06/common"
)

var (
	ErrEmailorPasswordInvalid = errors.New("email or password invalid")
	ErrEmailExisted           = errors.New("Email is Existed")
)

type User struct {
	common.SQLModel
	Email     string        `json:"email"gorm:"column:email;"`
	Password  string        `json:"password" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"last_name"gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"role"gorm:"column:role;"`
	Status    int           `json:"status"gorm:"column:status;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) GetUsedID() int {
	return u.ID
}
func (u *User) GetRole() string {
	return u.Role
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) Mask(isAdmin bool) {
	u.SQLModel.GenUID(common.DbTypeUser)
}
func (u User) TableName() string { return "users" }

type UserCreate struct {
	common.SQLModel
	Email     string        `json:"email"gorm:"column:email;"`
	Password  string        `json:"password" gorm:"column:password;"`
	Salt      string        `json:"-" gorm:"column:salt;"`
	LastName  string        `json:"last_name"gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"role"gorm:"column:role;"`
	Status    int           `json:"status"gorm:"column:status;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.SQLModel.GenUID(common.DbTypeUser)
}

func (u UserCreate) TableName() string { return User{}.TableName() }

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}
