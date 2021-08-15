package ginuser

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	"github.com/dachanh/food-delivery-G06/component/hasher"
	userbusiness "github.com/dachanh/food-delivery-G06/module/user/business"
	usermodel "github.com/dachanh/food-delivery-G06/module/user/model"
	userstorage "github.com/dachanh/food-delivery-G06/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(appContext appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appContext.GetMaiDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbusiness.NewUserRegister(store, md5)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSONP(http.StatusOK, map[string]string{
			"data": strconv.Itoa(data.ID),
		})
	}
}
