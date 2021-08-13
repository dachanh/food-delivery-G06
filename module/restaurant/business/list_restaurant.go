package businessrestaurant

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantbiz struct {
	store ListRestaurantStore
}

func NewListRestaurantbiz(store ListRestaurantStore) *listRestaurantbiz {
	return &listRestaurantbiz{
		store: store,
	}
}
func (biz *listRestaurantbiz) ListDataWithCondition(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, "User")
	if err != nil {
		return nil, err
	}
	return result, nil
}
