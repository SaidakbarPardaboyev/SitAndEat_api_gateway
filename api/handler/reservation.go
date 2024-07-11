package handler

import (
	pb "api_gateway/genproto/resirvation"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Yangi rezervatsiyani yaratish
// @Description So'rov tanasiga asoslangan holda yangi rezervatsiyani yaratish
// @Tags rezervatsiyalar
// @Accept json
// @Produce json
// @Param reservation body pb.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} pb.ResponseReservations
// @Failure 400 {object} gin.H
// @Router /reservations [post]
func(h *Handler) CreateReservation(c *gin.Context){
	req := pb.RequestReservations{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Create reservation error: %v", err)
		return 
	}

	resp, err := h.ReservationClient.Createreservations(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Create reservation request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}


// @Summary reservation olish
// @Description Id bilan reservation olinyapti
// @Tags rezervatsiyalar
// @Accept json
// @Produce json
// @Param reservation body pb.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} pb.ResponseReservations
// @Failure 400 {object} gin.H
// @Router /reservations [post]
func(h *Handler) GetReservation(c *gin.Context){
	id := c.Param("id")
	req := pb.ReservationId{Id: id}

	resp, err := h.ReservationClient.GetByIdReservations(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Get reservation request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}