package restaurantmodel

import (
	"encoding/json"
	"github.com/dachanh/food-delivery-G06/common"
)

type Restaurant struct {
	common.SQLModel
	OwnerID int             `json:"owner_id" gorm:"column:owner_id"`
	Name    string          `json:"name" gorm:"column:name;"`
	Address string          `json:"address" gorm:"column:addr;"`
	CityID  int             `json:"city_id" gorm:"column:city_id;"`
	lat     float64         `json:"lat" gorm:"column:lat;"`
	lng     float64         `json:"lng" gorm:"column:lng;"`
	cover   json.RawMessage `json:"cover" gorm:"column:cover;"`
	logo    json.RawMessage `json:"logo" gorm:"column:log;"`
}

func (Restaurant) TableName() string { return "restaurants" }
