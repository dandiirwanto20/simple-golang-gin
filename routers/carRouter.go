package routers

import (
	"gin-framework/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// create car
	router.POST("/cars", controllers.CreateCar)

	// update car
	router.PUT("/cars/:carID", controllers.UpdateCar)

	// get car
	router.GET("/cars/:carID", controllers.GetCar)

	// delete car
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
