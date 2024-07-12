package handler

import (
	pb "api_gateway/genproto/restaurant"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary restaurant olish
// @Description Id bilan restaurant olinyapti
// @Tags restaurant
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param restaurant body restaurant.RestuanantId true "Restaurant so'rovi"
// @Success 200 {object} restaurant.GetRes
// @Failure 400 {object} models.Error
// @Router /restaurant/getRestaurant/:id [get]
func (h *Handler) GetRestaurant(c *gin.Context) {
	req := pb.RestuanantId{
		Id: c.Param("id"),
	}

	resp, err := h.RestaurantClient.GetRestuarant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("GetRestuarant request error : %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary restaurant yangilash
// @Description Bodydagi ma'luotlar asosida restaurant yangilanyapti
// @Tags restaurant
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param restaurant body restaurant.RestuarantUpdate true "Restaurant so'rovi"
// @Success 200 {object} restaurant.Status
// @Failure 400 {object} models.Error
// @Router /restaurant/updateRestaurant [put]
func (h *Handler) UpdateRestaurant(c *gin.Context) {
	req := pb.RestuarantUpdate{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("Update restaurant error : %v", err)
		return
	}

	resp, err := h.RestaurantClient.UpdateRestuarant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("UpdateRestaurant request error : %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary restaurant o'chirish
// @Description Id orqali restaurant o'chirilmoqda
// @Tags restaurant
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param restaurant body restaurant.RestuanantId true "Restaurant so'rovi"
// @Success 200 {object} restaurant.Status
// @Failure 400 {object} models.Error
// @Router /restaurant/deleteRestaurant/:id [delete]
func (h *Handler) DeleteRestaurant(c *gin.Context) {
	req := pb.RestuanantId{
		Id: c.Param("id"),
	}

	resp, err := h.RestaurantClient.DeleteRestuarant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("DeleteRestuarant request error : %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary restaurant yaratish
// @Description Bodydagi ma'luotlar asosida restaurant yaratilyapti
// @Tags restaurant
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param restaurant body restaurant.Restuarant true "Restaurant so'rovi"
// @Success 200 {object} restaurant.Status
// @Failure 400 {object} models.Error
// @Router /restaurant/createRestaurant [post]
func (h *Handler) CreateRestaurant(c *gin.Context) {
	req := pb.Restuarant{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("CreateRestaurant  error : %v", err)
		return
	}

	resp, err := h.RestaurantClient.CreateRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("CreateRestaurant request error : %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Barcha restaurantni olish
// @Description Istalgancha restaurantni filterlab olish mumkin
// @Tags restaurant
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param restaurant body restaurant.FilterField true "Restaurant so'rovi"
// @Success 200 {object} restaurant.Restuanants
// @Failure 400 {object} models.Error
// @Router /restaurant/getAllRestaurant [get]
func (h *Handler) GetAllRestaurants(c *gin.Context) {
	req := pb.FilterField{
		Name:      c.Param("name"),
		Address:   c.Param("address"),
		CreatedAt: c.Param("createdAt"),
		Limit:     c.Param("limit"),
		Offset:    c.Param("offset"),
	}

	resp, err := h.RestaurantClient.GetAllRestaurants(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.Logger.Error("GetAll request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
