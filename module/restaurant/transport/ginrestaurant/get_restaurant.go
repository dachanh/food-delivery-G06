package ginrestaurant

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	businessrestaurant "github.com/dachanh/food-delivery-G06/module/restaurant/business"
	restaurantstorage "github.com/dachanh/food-delivery-G06/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message error": err.Error()})
			return
		}
		db := appContext.GetMaiDBConnection()

		store := restaurantstorage.NewSqlStore(db)

		getRestaurantBiz := businessrestaurant.NewGetRestaurantbiz(store)
		data, err := getRestaurantBiz.GetDataRestaurant(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message error": err.Error()})
			return
		}
		data.Mask(false)
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
