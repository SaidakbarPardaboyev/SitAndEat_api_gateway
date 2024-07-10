package api

import (
	"api_gateway/api/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)


func Router(conn *grpc.ClientConn)*gin.Engine{
	router := gin.Default()
	h := handler.NewHandlerRepo(conn)

	users := router.Group("/users")
	users.GET("/getProfile/:id", h.GetProfile)
	users.PUT("/updateProfile", h.UpdateProfile)
	users.DELETE("/deleteProfile", h.DeleteProfile)
}