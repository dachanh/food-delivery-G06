package restaurantlikesstorage

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantlikesmodel "github.com/dachanh/food-delivery-G06/module/restaurantlikes/model"
)

func (s *sqlStore) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)
	type sqlData struct {
		RestaurantID int `gorm:"column:restaurnt_id;"`
		LikeCount    int `gorm:"column:count;"`
	}
	var listLike []sqlData
	if err := s.db.Table(restaurantlikesmodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	for _, item := range listLike {
		result[item.RestaurantID] = item.LikeCount
	}
	return result, nil
}
