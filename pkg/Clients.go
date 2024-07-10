package pkg

import (
	ppay "api_gateway/genproto/payment"
	presirvation "api_gateway/genproto/resirvation"
	pmenu "api_gateway/genproto/menu"
	prestaurant "api_gateway/genproto/restaurant"
	"api_gateway/config"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPaymentClient(cfg *config.Config) ppay.PaymentClient {
	conn, err := grpc.NewClient(cfg.PAYMENT_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return ppay.NewPaymentClient(conn)
}

func NewReservationClient(cfg *config.Config) presirvation.ResirvationClient {
	conn, err := grpc.NewClient(cfg.RESERVATION_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return presirvation.NewResirvationClient(conn)
}

func NewRestaurantClient(cfg *config.Config) prestaurant.RestaurantClient {
	conn, err := grpc.NewClient(cfg.RESERVATION_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return prestaurant.NewRestaurantClient(conn)
}

func NewMenuClient(cfg *config.Config) pmenu.MenuClient {
	conn, err := grpc.NewClient(cfg.RESERVATION_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return pmenu.NewMenuClient(conn)
}