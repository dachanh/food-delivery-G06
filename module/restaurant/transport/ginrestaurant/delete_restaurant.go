package ginrestaurant

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	businessrestaurant "github.com/dachanh/food-delivery-G06/module/restaurant/business"
	restaurantstorage "github.com/dachanh/food-delivery-G06/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := restaurantstorage.NewSqlStore(appContext.GetMaiDBConnection())

		deleteRestaurantbiz := businessrestaurant.NewDeleteRestaurantBiz(store)
		if err := deleteRestaurantbiz.DeleteRestaurant(c.Request.Context(), id, true); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": 1})
	}
}
