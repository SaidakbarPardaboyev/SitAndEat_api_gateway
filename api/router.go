package api

import (
	"api_gateway/api/handler"
	"api_gateway/config"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Router(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	h := handler.NewHandlerRepo(config.Load())

	users := router.Group("/users")
	users.GET("/getProfile/:id", h.GetProfile)
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

	return router
}
