package restaurantmodel

import (
	"encoding/json"
	"github.com/dachanh/food-delivery-G06/common"
)

type RestaurantCreate struct {
	common.SQLModel
	OwnerID          int             `json:"owner_id" gorm:"column:owner_id"`
	Name             string          `json:"name" gorm:"column:name;"`
	Address          string          `json:"address" gorm:"column:addr;"`
	CityID           int             `json:"city_id" gorm:"column:city_id;"`
	ShippingFeePerKM float64         `json:"shipping_fee_per_km" gorm:"column:"shipping_fee_per_km";`
	Lat              float64         `json:"lat" gorm:"column:lat;"`
	Lng              float64         `json:"lng" gorm:"column:lng;"`
	Cover            json.RawMessage `json:"cover" gorm:"column:cover;"`
	Logo             json.RawMessage `json:"logo" gorm:"column:logo;"`
	//Cover            json.RawMessage `json:"cover" gorm:"column:cover;"`
	//Logo             json.RawMessage `json:"logo" gorm:"column:log;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }
