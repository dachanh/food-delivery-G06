package restaurantlikebusiness

import (
	"context"
	restaurantlikesmodel "github.com/dachanh/food-delivery-G06/module/restaurantlikes/model"
)

type RestaurantDislikeStore interface {
	DeleteUserDislikeRestaurant(ctx context.Context, restaurantId int, userId int) error
}

type restaurantDislikeBusiness struct {
	store RestaurantDislikeStore
}

func NewRestaurantDislikeBiz(store RestaurantDislikeStore) *restaurantDislikeBusiness {
	return &restaurantDislikeBusiness{
		store: store,
	}
}

func (biz *restaurantDislikeBusiness) UserDislikeRestaurant(ctx context.Context, restaurantId int, userId int) error {
	err := biz.store.DeleteUserDislikeRestaurant(ctx, restaurantId, userId)
	if err != nil {
		return restaurantlikesmodel.ErrCannotDislikeRestaurant(err)
	}
	return nil
}
