package handler

import (
	pb "api_gateway/genproto/resirvation"
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

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
func (h *Handler) CreateReservation(c *gin.Context) {
	req := pb.RequestReservations{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Create reservation error: %v", err)
		return
	}

	resp, err := h.ReservationClient.Createreservations(context.Background(), &req)
	if err != nil {
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
func (h *Handler) GetReservation(c *gin.Context) {
	id := c.Param("id")
	req := pb.ReservationId{Id: id}

	resp, err := h.ReservationClient.GetByIdReservations(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Get reservation request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary barcha reservationlarni olish
// @Description istalgan reservatinlarni filterlab olish mumkin
// @Tags rezervatsiyalar
// @Accept json
// @Produce json
// @Param reservation body pb.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} pb.ResponseReservations
// @Failure 400 {object} gin.H
// @Router /reservations [post]
func (h *Handler) GetAllReservation(c *gin.Context) {
	query := `
			SELECT 
				* 
			FROM 
				Reservation 
			WHERE 
				deleted_at is null`
	param := []string{}
	arr := []string{}

	req := pb.Reservation{
		Id:           c.Query("id"),
		UserId:       c.Query("userId"),
		RestuarantId: c.Query("restaurantId"),
		ResTime:      c.Query("resTime"),
		Status:       c.Query("status"),
		CreatedAt:    c.Query("createdAt"),
		UpdateAt:     c.Query("updatedAt"),
	}
	limit := c.Query("limit")
	offset := c.Query("offset")

	if len(req.Id) > 0 {
		query += " and id = :id"
		param = append(param, ":id")
		arr = append(arr, req.Id)
	}
	if len(req.UserId) > 0 {
		query += " and user_id = :user_id"
		param = append(param, ":user_id")
		arr = append(arr, req.UserId)
	}

	if len(req.RestuarantId) > 0 {
		query += " and restaurant_id = :restaurant_id"
		param = append(param, ":restaurant_id")
		arr = append(arr, req.RestuarantId)
	}

	if len(req.ResTime) > 0 {
		query += " and reservation_time = :reservation_time"
		param = append(param, ":reservation_time")
		arr = append(arr, req.ResTime)
	}

	if len(req.Status) > 0 {
		query += " and status = :status"
		param = append(param, ":status")
		arr = append(arr, req.Status)
	}

	if len(req.CreatedAt) > 0 {
		data := strings.Split(req.CreatedAt, "-")
		query += " and created_at BETWEEN :created_at1 and :created_at2"
		param = append(param, ":created_at1", ":created_at2")
		arr = append(arr, data[0], data[1])
	}

	if len(req.UpdateAt) > 0 {
		data := strings.Split(req.CreatedAt, "-")
		query += " and updated_at BETWEEN :updated_at1 and :updated_at2"
		param = append(param, ":updated_at1", ":updated_at2")
		arr = append(arr, data[0], data[1])
	}

	if len(limit) > 0 {
		query += " limit " + limit
	}

	if len(offset) > 0 {
		query += " offset " + offset
	}

	for k, v := range param {
		query = strings.Replace(query, v, "$"+strconv.Itoa(k), 1)
	}

	resp, err := h.ReservationClient.GetAllReservations(context.Background(), &pb.Filter{Query: query, Arr: arr})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("GetAll request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary reservationni yangilash
// @Description bodydan yangi reservationni olgan holda uni yangilamoqda
// @Tags rezervatsiyalar
// @Accept json
// @Produce json
// @Param reservation body pb.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} pb.ResponseReservations
// @Failure 400 {object} gin.H
// @Router /reservations [post]
func(h *Handler) UpdateReservations(c *gin.Context){
	req := pb.ReservationUpdate{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Update reservation error %v: ", err)
		return
	}

	resp, err := h.ReservationClient.UpdateReservations(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("UpdateReservation request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary reservationni o'chirish
// @Description reservation id sini paramdan o'qigan holda uni o'chirmoqda
// @Tags rezervatsiyalar
// @Accept json
// @Produce json
// @Param reservation body pb.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} pb.ResponseReservations
// @Failure 400 {object} gin.H
// @Router /reservations [post]
func(h *Handler) DeleteReservation(c *gin.Context){
	req := pb.ReservationId{
		Id: c.Param("id"),
	}

	resp, err := h.ReservationClient.DeleteReservations(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("DeleteReservation request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary user reservationni olish
// @Description user id sini paramdan o'qigan holda uning reservationi olinmoqda
// @Tags rezervatsiyalar
// @Accept json
// @Produce json
// @Param reservation body pb.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} pb.ResponseReservations
// @Failure 400 {object} gin.H
// @Router /reservations [post]
func(h *Handler) GetReservationsByUserId(c *gin.Context){
	req := pb.UserId{
		Id: c.Param("userId"),
	}

	resp, err := h.ReservationClient.GetReservationsByUserId(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Get userReservation request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary order meal
// @Description bodydagi ma'lumotlar asosida ovqat zakaz qilinmoqda
// @Tags rezervatsiyalar
// @Accept json
// @Produce json
// @Param reservation body pb.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} pb.ResponseReservations
// @Failure 400 {object} gin.H
// @Router /reservations [post]
func(h *Handler) OrderMeal(c *gin.Context){
	req := pb.Order{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Order Meal error: %v", err)
		return
	}

	resp, err := h.ReservationClient.OrderMeal(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("OrderMeal request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary pay for reservation
// @Description bodydagi ma'lumotlar asosida to'lob qilinishi kerak
// @Tags rezervatsiyalar
// @Accept json
// @Produce json
// @Param reservation body pb.RequestReservations true "Rezervatsiya so'rovi"
// @Success 200 {object} pb.ResponseReservations
// @Failure 400 {object} gin.H
// @Router /reservations [post]
func(h *Handler) PayForReservation(c *gin.Context){
	req := pb.Payment{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("Payment for reservation error: %v", err)
		return
	}

	resp, err := h.ReservationClient.PayForReservation(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		log.Fatalf("PayForReservation request error %v: ", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}







