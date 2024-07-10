package handler

import (
	"api_gateway/genproto/menu"
	"api_gateway/genproto/payment"
	"api_gateway/genproto/resirvation"
	"api_gateway/genproto/restaurant"
	"api_gateway/genproto/users"

	"google.golang.org/grpc"
)


type Handler struct{
	users.UsersClient
	menu.MenuClient
	payment.PaymentClient
	resirvation.ResirvationClient
	restaurant.RestaurantClient
}

func NewHandlerRepo(Conn *grpc.ClientConn)*Handler{
	return &Handler{
		users.NewUsersClient(Conn),
		menu.NewMenuClient(Conn),
		payment.NewPaymentClient(Conn),
		resirvation.NewResirvationClient(Conn),
		restaurant.NewRestaurantClient(Conn),	
	}
}