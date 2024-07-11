package handler

import (
	pb "api_gateway/genproto/menu"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get menu
// @Description get menu
// @Tags menu
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetMenuResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu [get]
func (h *Handler) CreateFood(c *gin.Context) {
	req := &pb.CreateF{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.MenuClient.CreateFood(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get All foods
// @Description get all foods
// @Tags menu
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetMenuResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu [get]
func (h *Handler) GetAllFoods(c *gin.Context) {
	req := &pb.Void{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.MenuClient.GetAllFoods(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Summary Get food
// @Description get food
// @Tags menu
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetMenuResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu/:id [get]
func (h *Handler) GetFood(c *gin.Context) {
	req := &pb.FoodId{}

	req.Id = c.Param("id")
	resp, err := h.MenuClient.GetFood(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Update food
// @Description update food
// @Tags menu
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetMenuResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu/:id [put]
func (h *Handler) UpdateFood(c *gin.Context) {
	req := &pb.UpdateF{}
	resp, err := h.MenuClient.UpdateFood(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete food
// @Description delete food
// @Tags menu
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetMenuResponse
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu/:id [delete]
func (h *Handler) DeleteFood(c *gin.Context) {
	req := &pb.FoodId{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := c.Param("id")
	req.Id = id	
	resp, err := h.MenuClient.DeleteFood(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}