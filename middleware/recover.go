package middleware

import (
	"github.com/dachanh/food-delivery-G06/common"
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/gin-gonic/gin"
)

func Recover(app appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				return

			}
		}()
		c.Next()
	}
}
