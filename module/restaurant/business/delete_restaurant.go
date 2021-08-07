package businessrestaurant

import (
	"context"
	"errors"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type DeleteResturantStore interface {
	Delete(cxt context.Context, id int) error
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
	Update(ctx context.Context, id int, update *restaurantmodel.RestaurantUpdate) error
}

type deleteRestaurantBiz struct {
	store DeleteResturantStore
}

func NewDeleteRestaurantBiz(store DeleteResturantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store: store,
	}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int, isSoft bool) error {
	oldData, err := biz.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrDB(err)
	}
	if oldData.Status == 0 {
		return common.ErrDB(errors.New("restaurant has been deleted"))
	}
	if isSoft {
		zero := 0
		if err := biz.store.Update(ctx, id, &restaurantmodel.RestaurantUpdate{Status: &zero}); err != nil {
			return common.ErrDB(err)
		}
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrDB(err)
	}
	return nil
}
