package restaurantlikesstorage

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantlikesmodel "github.com/dachanh/food-delivery-G06/module/restaurantlikes/model"
	"gorm.io/gorm"
)

func (s *sqlStore) DeleteUserDislikeRestaurant(ctx context.Context, restaurantId int, userId int) error {
	db := s.db.Begin()
	var tempRestaurant restaurantlikesmodel.Like
	if err := db.Table(restaurantlikesmodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id =?", userId, restaurantId).
		First(&tempRestaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrDB(common.RecordNotFound)
		}
		return common.ErrDB(err)
	}
	if err := db.Table(restaurantlikesmodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id =?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
