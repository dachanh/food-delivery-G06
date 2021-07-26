package ginrestaurant

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	businessrestaurant "github.com/dachanh/food-delivery-G06/module/restaurant/business"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
	restaurantstorage "github.com/dachanh/food-delivery-G06/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var newRestaurant restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&newRestaurant); err != nil {
			//log.Fatal(err)
			c.JSONP(http.StatusBadRequest, gin.H{"message error": err.Error()})
			return
		}
		// store layer
		store := restaurantstorage.NewSqlStore(appContext.GetMaiDBConnection())
		// business layer
		biz := businessrestaurant.NewCreateRestaurantBiz(store)
		err := biz.CreateRestaurant(c.Request.Context(), &newRestaurant)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": newRestaurant.ID})
	}
}
