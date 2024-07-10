package handler

import (
	"api_gateway/genproto/menu"
	"api_gateway/genproto/payment"
	"api_gateway/genproto/resirvation"
	"api_gateway/genproto/restaurant"
	"api_gateway/genproto/users"

	"google.golang.org/grpc"
)

type Handler struct {
	UserClient        users.UsersClient
	MenuClient        menu.MenuClient
	PaymentClient     payment.PaymentClient
	ReservationClient resirvation.ResirvationClient
	RestaurantClient  restaurant.RestaurantClient
}

func NewHandlerRepo(Conn *grpc.ClientConn) *Handler {
	return &Handler{
		UserClient:        users.NewUsersClient(Conn),
		MenuClient:        menu.NewMenuClient(Conn),
		PaymentClient:     payment.NewPaymentClient(Conn),
		ReservationClient: resirvation.NewResirvationClient(Conn),
		RestaurantClient:  restaurant.NewRestaurantClient(Conn),
	}
}
