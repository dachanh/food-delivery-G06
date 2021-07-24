package businessrestaurant

import (
	"context"
	"errors"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type CreateStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateStore
}

func NewCreateRestaurantBiz(store CreateStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("valid name field")
	}
	err := biz.store.Create(ctx, data)
	return err
}
