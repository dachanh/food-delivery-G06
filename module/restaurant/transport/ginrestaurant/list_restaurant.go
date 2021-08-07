package ginrestaurant

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": 1})
	}
}
