package ginrestaurantlike

import (
	"github.com/dachanh/food-delivery-G06/common"
	"github.com/dachanh/food-delivery-G06/component/appctx"
	restaurantlikebusiness "github.com/dachanh/food-delivery-G06/module/restaurantlikes/business"
	restaurantlikesstorage "github.com/dachanh/food-delivery-G06/module/restaurantlikes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestaurantDislike(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := restaurantlikesstorage.NewSqlStore(appCtx.GetMaiDBConnection())
		biz := restaurantlikebusiness.NewRestaurantDislikeBiz(store)
		err = biz.UserDislikeRestaurant(c.Request.Context(), int(uid.GetLocalID()), requester.GetUsedID())

		if err != nil {
			panic(err)
		}
		c.JSONP(http.StatusOK, map[string]string{
			"message": "ok",
		})
	}
}
