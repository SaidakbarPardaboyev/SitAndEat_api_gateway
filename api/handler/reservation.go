package handler

import (
	pb "api_gateway/genproto/resirvation"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Yangi rezervatsiyani yaratish
// @Description So'rov tanasiga asoslangan holda yangi rezervatsiyani yaratish
// @Tags resirvation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param reservation body resirvation.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} resirvation.Status
// @Failure 400 {object} models.Error
// @Router /reservation/createReservation [post]
func (h *Handler) CreateReservation(c *gin.Context) {
	req := pb.RequestReservations{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("Create reservation error: %v", err)
		return
	}

	resp, err := h.ReservationClient.Createreservations(context.Background(), &req)
	if err != nil {
		fmt.Println(err)
		h.Logger.Error("Create reservation request error: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary reservation olish
// @Description Id bilan reservation olinyapti
// @Tags resirvation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Reservation ID"
// @Success 200 {object} resirvation.Reservation
// @Failure 400 {object} models.Error
// @Router /reservation/getReservation/{id} [get]
func (h *Handler) GetReservation(c *gin.Context) {
	id := c.Param("id")
	req := pb.ReservationId{Id: id}

	resp, err := h.ReservationClient.GetByIdReservations(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("Get reservation request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Barcha reservationlarni olish
// @Description Istalgan reservatinlarni filterlab olish mumkin
// @Tags reservation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param status query string false "Status of the reservations"
// @Param createdAt query string false "Creation date of the reservations"
// @Param updatedAt query string false "Last update date of the reservations"
// @Param limit query integer false "Limit the number of results"
// @Param offset query integer false "Offset for pagination"
// @Success 200 {object} resirvation.Reservations
// @Failure 400 {object} models.Error
// @Router /reservation/getAllReservations [get]
func (h *Handler) GetAllReservation(c *gin.Context) {
	req := pb.FilterField{
		Status:    c.Query("status"),
		CreatedAt: c.Query("createdAt"),
		UpdateAt:  c.Query("updatedAt"),
		Limit:     c.Query("limit"),
		Offset:    c.Query("offset"),
	}

	resp, err := h.ReservationClient.GetAllReservations(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("GetAll request error: %v", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary reservationni yangilash
// @Description bodydan yangi reservationni olgan holda uni yangilamoqda
// @Tags resirvation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param reservation body resirvation.ReservationUpdate true "Rezervatsiya so'rovi"
// @Success 200 {object} resirvation.Status
// @Failure 400 {object} models.Error
// @Router /reservation/updateReservation [put]
func (h *Handler) UpdateReservations(c *gin.Context) {
	req := pb.ReservationUpdate{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("Update reservation error %v: ", err)
		return
	}

	resp, err := h.ReservationClient.UpdateReservations(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("UpdateReservation request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary reservationni o'chirish
// @Description reservation id sini paramdan o'qigan holda uni o'chirmoqda
// @Tags resirvation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Reservation ID"
// @Success 200 {object} resirvation.Status
// @Failure 400 {object} models.Error
// @Router /reservation/deleteReservation/{id} [delete]
func (h *Handler) DeleteReservation(c *gin.Context) {
	req := pb.ReservationId{
		Id: c.Param("id"),
	}

	resp, err := h.ReservationClient.DeleteReservations(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("DeleteReservation request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary user reservationni olish
// @Description user id sini paramdan o'qigan holda uning reservationi olinmoqda
// @Tags resirvation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} resirvation.Reservations
// @Failure 400 {object} models.Error
// @Router /reservation/getUserReservation/{id} [get]
func (h *Handler) GetReservationsByUserId(c *gin.Context) {
	req := pb.UserId{
		Id: c.Param("id"),
	}
	fmt.Println(req.Id)

	resp, err := h.ReservationClient.GetReservationsByUserId(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("Get userReservation request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary order meal
// @Description bodydagi ma'lumotlar asosida ovqat zakaz qilinmoqda
// @Tags resirvation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param reservation body resirvation.Order true "Rezervatsiya so'rovi"
// @Success 200 {object} resirvation.Status
// @Failure 400 {object} models.Error
// @Router /reservation/orderMeal [post]
func (h *Handler) OrderMeal(c *gin.Context) {
	req := pb.Order{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("Order Meal error: %v", err)
		return
	}

	resp, err := h.ReservationClient.OrderMeal(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("OrderMeal request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary pay for reservation
// @Description bodydagi ma'lumotlar asosida to'lob qilinishi kerak
// @Tags resirvation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param reservation body resirvation.Payment true "Rezervatsiya so'rovi"
// @Success 200 {object} resirvation.Status
// @Failure 400 {object} models.Error
// @Router /reservation/payForReservation [post]
func (h *Handler) PayForReservation(c *gin.Context) {
	req := pb.Payment{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("Payment for reservation error: %v", err)
		return
	}

	resp, err := h.ReservationClient.PayForReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("PayForReservation request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
