package ginrestaurantlike

import (
	"github.com/dachanh/food-delivery-G06/common"
	"github.com/dachanh/food-delivery-G06/component/appctx"
	restaurantlikebusiness "github.com/dachanh/food-delivery-G06/module/restaurantlikes/business"
	restaurantlikesmodel "github.com/dachanh/food-delivery-G06/module/restaurantlikes/model"
	restaurantlikesstorage "github.com/dachanh/food-delivery-G06/module/restaurantlikes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestaurantLike(appContext appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		likedata := restaurantlikesmodel.Like{
			RestaurantID: int(uid.GetLocalID()),
			UserID:       requester.GetUsedID(),
		}
		store := restaurantlikesstorage.NewSqlStore(appContext.GetMaiDBConnection())
		biz := restaurantlikebusiness.NewRestaurantlikeBiz(store)
		err = biz.RestaurantLike(c.Request.Context(), &likedata)
		if err != nil {
			panic(err)
		}
		c.JSONP(http.StatusOK, map[string]string{
			"message": "ok",
		})
	}
}
