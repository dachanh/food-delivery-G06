package restaurantlikesstorage

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantlikesmodel "github.com/dachanh/food-delivery-G06/module/restaurantlikes/model"
)

func (s *sqlStore) CreateRestaurantLikes(ctx context.Context, data *restaurantlikesmodel.Like) error {
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
