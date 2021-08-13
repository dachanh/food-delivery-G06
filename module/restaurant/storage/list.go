package restaurantstorage

import (
	"context"
	"github.com/dachanh/food-delivery-G06/common"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db
	db = db.Where("status in (?)", 1)
	if filter.OwnerId > 0 {
		db.Where("owner_id = ?", filter.OwnerId)
	}
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	for i := range moreKeys {
		if moreKeys[i] == "User" {
			db = db.Preload("User")
		}
	}
	if err := db.Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
