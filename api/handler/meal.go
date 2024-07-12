package handler

import (
	pb "api_gateway/genproto/menuRedis"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Yangi meal yaratish
// @Description Yangi meal yaratish uchun endpoint
// @Tags meal
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param meal body menuRedis.MealCreate true "Yangi meal so'rovi"
// @Success 200 {object} menuRedis.Status
// @Failure 400 {object} models.Error
// @Router /meal/createMeal [post]
func (h *Handler) CretaeMeal(c *gin.Context) {
	req := pb.MealCreate{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Meal create qilishda xato: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.Redis.CretaeMeal(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("MealCreate request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Mealni yangilash
// @Description Mealni yangilash
// @Tags meal
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param meal body menuRedis.MealCreate true "Yangi meal so'rovi"
// @Success 200 {object} menuRedis.Status
// @Failure 400 {object} models.Error
// @Router /meal/updateMeal [put]
func (h *Handler) UpdateMeal(c *gin.Context) {
	req := pb.MealCreate{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Meal update qilishda xato: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.Redis.UpdateMeal(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("MealUpdate request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Mealni o'chirish
// @Description Mealni o'chirish
// @Tags meal
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param meal body menuRedis.MealDelete true "Yangi meal so'rovi"
// @Success 200 {object} menuRedis.Status
// @Failure 400 {object} models.Error
// @Router /meal/deleteMeal [delete]
func (h *Handler) DeleteMeal(c *gin.Context) {
	req := pb.MealDelete{
		MealId: c.Param("id"),
	}

	resp, err := h.Redis.DeleteMeal(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("MealUpdate request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
