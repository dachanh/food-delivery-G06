package main

import (
	"github.com/dachanh/food-delivery-G06/component/appctx"
	ginrestaurant "github.com/dachanh/food-delivery-G06/module/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func Activate() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return err
	}
	dsn := os.Getenv("DNS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	appContext := appctx.NewAppContext(db)
	route := gin.Default()
	v1 := route.Group("/v1")
	{
		restaurant := v1.Group("/restaurant")
		{
			restaurant.POST("", ginrestaurant.CreateRestaurant(appContext))
		}
	}
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	route.Run()
	return nil
}

func main() {
	if err := Activate(); err != nil {
		log.Fatal(err)
	}
}
