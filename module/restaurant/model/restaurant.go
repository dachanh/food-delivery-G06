package restaurantmodel

import (
	"github.com/dachanh/food-delivery-G06/common"
)

type Restaurant struct {
	common.SQLModel
	Name    string        `json:"name" gorm:"column:name;"`
	Address string        `json:"address" gorm:"column:addr;"`
	Lat     float64       `json:"lat" gorm:"column:lat;"`
	Lng     float64       `json:"lng" gorm:"column:lng;"`
	Cover   *common.Image `json:"cover" gorm:"column:cover;"`
	Logo    *common.Image `json:"logo" gorm:"column:logo;"`
}

func (Restaurant) TableName() string { return "restaurants" }
