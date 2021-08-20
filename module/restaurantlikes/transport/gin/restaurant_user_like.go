package ginrestaurantlike

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/gin-gonic/gin"
)

func ListUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//uid, err := common.FromBase58(c.Param("id"))
		//if err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		//filter := restaurantlikesmodel.Filter{
		//	RestaurantID: int(uid.GetLocalID()),
		//}
		//var paging common.Paging
		//
		//if err := c.ShouldBind(&paging); err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		//paging.Process()
		//store := restaurantstorage.NewSqlStore(appCtx.GetMaiDBConnection())
		//likeStore := restaurantlikesstorage.NewSqlStore(appCtx.GetMaiDBConnection())
		//result, err := biz

	}
}
