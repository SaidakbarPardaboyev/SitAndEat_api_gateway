package api

import (
	"api_gateway/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description Bu API Gatewayning Swagger hujjatlari
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func NewRouter() *gin.Engine {
	router := gin.Default()
	h := handler.NewHandlerRepo()

	users := router.Group("/users")
	users.GET("/getProfile/:id", h.GetProfileById)
	users.PUT("/updateProfile", h.UpdateProfile)
	users.DELETE("/deleteProfile", h.DeleteProfile)

	restaurant := router.Group("/restaurant")
	restaurant.GET("/getRestaurant/:id", h.GetRestaurant)
	restaurant.PUT("/updateRestaurant", h.UpdateRestaurant)
	restaurant.DELETE("/deleteRestaurant/:id", h.DeleteRestaurant)
	restaurant.POST("/createRestaurant", h.CreateRestaurant)
	restaurant.GET("/getAllRestaurant", h.GetAllRestaurants)

	reservation := router.Group("/reservation")
	reservation.POST("/createReservation", h.CreateReservation)
	reservation.GET("/getReservation/:id", h.GetReservation)
	reservation.GET("/getAllReservations", h.GetAllReservation)
	reservation.PUT("/updateReservation", h.UpdateReservations)
	reservation.DELETE("/deleteReservation/:id", h.DeleteReservation)
	reservation.GET("/getUserReservation/:id", h.GetReservationsByUserId)
	reservation.POST("/orderMeal", h.OrderMeal)
	reservation.POST("/payForReservation", h.PayForReservation)

	menu := router.Group("/menu")
	menu.POST("/createFood", h.CreateFood)
	menu.GET("/getAllFoods", h.GetAllFoods)
	menu.GET("/getFood/:id", h.GetFood)
	menu.PUT("/updateFood", h.UpdateFood)
	menu.DELETE("/deleteFood/:id", h.DeleteFood)

	payment := router.Group("/payment")
	payment.POST("/createPayment", h.CreatePayments)
	payment.GET("/getPaymentStatus/:id", h.GetPaymentStatusById)
	payment.PUT("/updatePayment", h.UpdatePayments)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
