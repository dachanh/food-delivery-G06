package businessrestaurant

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type listRestaurantStore interface {
	ListDataWithCondition(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type ListRestaurantBiz struct {
	store listRestaurantStore
}

func (biz *ListRestaurantBiz) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
