package ginrestaurant

import (
	"github.com/dachanh/food-delivery-G06/common"
	"github.com/dachanh/food-delivery-G06/component/appctx"
	businessrestaurant "github.com/dachanh/food-delivery-G06/module/restaurant/business"
	restaurantmodel "github.com/dachanh/food-delivery-G06/module/restaurant/model"
	restaurantstorage "github.com/dachanh/food-delivery-G06/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db := restaurantstorage.NewSqlStore(appContext.GetMaiDBConnection())
		store := businessrestaurant.NewListRestaurantbiz(db)
		result, err := store.ListDataWithCondition(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for it := range result {
			result[it].Mask(false)
		}
		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging, "filter": filter})
	}
}
