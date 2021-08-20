package repositoryrestaurant

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

type GetLikedCountStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantLikeCountRepo struct {
	restaurantstore ListRestaurantStore
	likedstore      GetLikedCountStore
}

func NewListRestaurntLikeCountRepo(restaurantstore ListRestaurantStore, likedstore GetLikedCountStore) *listRestaurantLikeCountRepo {
	return &listRestaurantLikeCountRepo{
		restaurantstore: restaurantstore,
		likedstore:      likedstore,
	}
}

func (biz *listRestaurantLikeCountRepo) GetRestaurantLikeList(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	result, err := biz.restaurantstore.ListDataWithCondition(ctx, filter, paging, moreKeys...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
