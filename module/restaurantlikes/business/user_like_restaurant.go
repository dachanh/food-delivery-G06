package restaurantlikebusiness

import (
	"context"
	restaurantlikesmodel "github.com/dachanh/food-delivery-G06/module/restaurantlikes/model"
)

type RestaurantlikeStore interface {
	CreateUserLikeRestaurant(ctx context.Context, data *restaurantlikesmodel.Like) error
}

type restaurantlikebiz struct {
	store RestaurantlikeStore
}

func NewRestaurantlikeBiz(store RestaurantlikeStore) *restaurantlikebiz {
	return &restaurantlikebiz{
		store: store,
	}
}

func (biz *restaurantlikebiz) RestaurantLike(ctx context.Context, data *restaurantlikesmodel.Like) error {
	err := biz.store.CreateUserLikeRestaurant(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
