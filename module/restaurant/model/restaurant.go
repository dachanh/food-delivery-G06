package restaurantmodel

import (
	"github.com/dachanh/food-delivery-G06/common"
)

type Restaurant struct {
	common.SQLModel
	Name             string             `json:"name" gorm:"column:name;"`
	OwnerID          int                `json:"owner_id" gorm:"column:owner_id""`
	CityID           int                `json:"city_id"gorm:"city_id"`
	Address          string             `json:"address" gorm:"column:addr;"`
	Lat              float64            `json:"lat" gorm:"column:lat;"`
	Lng              float64            `json:"lng" gorm:"column:lng;"`
	shippingFeePerKm float64            `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km"`
	Cover            *common.Image      `json:"cover" gorm:"column:cover;"`
	Logo             *common.Image      `json:"logo" gorm:"column:logo;"`
	User             *common.SimpleUser `json:"user" gorm:"foreignKey:OwnerID;preload:false;"`
}

func (Restaurant) TableName() string { return "restaurants" }
func (data *Restaurant) Mask(isOwnerOrAdmin bool) {
	data.GenUID(common.DbTypeRestaurant)

	if u := data.User; u != nil {
		u.Mask(isOwnerOrAdmin)
	}
}
