package businessrestaurant

import (
	"context"
	"errors"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type GetRestaurantStore interface {
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantbiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{
		store: store,
	}
}

func (biz *getRestaurantBiz) GetDataRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrDB(err)
	}
	if data.Status == 0 {
		return nil, common.ErrDB(errors.New("restaurant has been deleted"))
	}
	return data, nil
}
