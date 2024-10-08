package pkg

import (
	"api_gateway/config"
	pmenu "api_gateway/genproto/menu"
	"api_gateway/genproto/menuRedis"
	ppay "api_gateway/genproto/payment"
	presirvation "api_gateway/genproto/resirvation"
	prestaurant "api_gateway/genproto/restaurant"
	user "api_gateway/genproto/users"
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

func NewUsersClient(cfg *config.Config) user.UsersClient {
	conn, err := grpc.NewClient(cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return user.NewUsersClient(conn)
}

func NewMenuRedisClient(cfg *config.Config) menuRedis.MenuClient {
	conn, err := grpc.NewClient(cfg.RESERVATION_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal(err)
	}
	return menuRedis.NewMenuClient(conn)
}