package restaurantmodel

import "github.com/dachanh/food-delivery-G06/common"

type RestaurantUpdate struct {
	common.SQLModel
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
	Status  *int    `json:"_" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
