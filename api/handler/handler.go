package handler

import (
	"api_gateway/config"
	"api_gateway/genproto/menu"
	"api_gateway/genproto/payment"
	"api_gateway/genproto/resirvation"
	"api_gateway/genproto/restaurant"
	"api_gateway/genproto/users"
	"api_gateway/pkg"
)

type Handler struct {
	UserClient        users.UsersClient
	MenuClient        menu.MenuClient
	PaymentClient     payment.PaymentClient
	ReservationClient resirvation.ResirvationClient
	RestaurantClient  restaurant.RestaurantClient
}

func NewHandlerRepo() *Handler {
	cfg := config.Load()
	return &Handler{
		UserClient:        pkg.NewUsersClient(&cfg),
		MenuClient:        pkg.NewMenuClient(&cfg),
		PaymentClient:     pkg.NewPaymentClient(&cfg),
		ReservationClient: pkg.NewReservationClient(&cfg),
		RestaurantClient:  pkg.NewRestaurantClient(&cfg),
	}
}
