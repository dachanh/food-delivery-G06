package businessrestaurant

import (
	"context"
	"errors"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
	Update(ctx context.Context, id int, updateData *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantbiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{
		store: store,
	}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := biz.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrDB(errors.New("error database"))
	}
	if oldData.Status == 0 {
		return common.ErrDB(errors.New("restaurant has been deleted"))
	}
	err = biz.store.Update(ctx, id, data)
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
