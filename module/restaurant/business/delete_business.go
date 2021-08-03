package businessrestaurant

import (
	"context"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type DeleteStore interface {
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
	Delete(cxt context.Context, id int) error
	GetDataWithConditions(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
}
