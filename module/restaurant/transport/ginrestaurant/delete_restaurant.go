package ginrestaurant

import (
	"github.com/dachanh/food-delivery-G06/common"
	"github.com/dachanh/food-delivery-G06/component/appctx"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
	restaurantstorage "github.com/dachanh/food-delivery-G06/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteRestaurant(appContext appctx.AppContext) func(c *gin.HandlerFunc) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var dataUpdate restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&dataUpdate); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSqlStore(appContext.GetMaiDBConnection())

		binz :=
	}
}
