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
// @Param request body menu.CreateF true "Create Food"
// @Success 200 {object} menu.Status
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu/createFood [post]
func (h *Handler) CreateFood(c *gin.Context) {
	req := &pb.CreateF{}
	if err := c.ShouldBindJSON(req); err != nil {
		h.Logger.Error("CreateFood error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.MenuClient.CreateFood(context.Background(), req)
	if err != nil {
		h.Logger.Error("CreateFood request error: %v", err)
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
// @Param request body menu.Void true "Get Foods"
// @Success 200 {object} menu.Foods
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu/getAllFoods [get]
func (h *Handler) GetAllFoods(c *gin.Context) {
	req := &pb.Void{}

	resp, err := h.MenuClient.GetAllFoods(context.Background(), req)
	if err != nil {
		h.Logger.Error("GetAllFoods request error: %v", err)
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
// @Param request body menu.FoodId true "Get Food"
// @Success 200 {object} menu.Food
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu/getFood/:id [get]
func (h *Handler) GetFood(c *gin.Context) {
	req := &pb.FoodId{
		Id: c.Param("id"),
	}
	
	resp, err := h.MenuClient.GetFood(context.Background(), req)
	if err != nil {
		h.Logger.Error("GetFood request error: %v", err)
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
// @Param request body menu.UpdateF true "Update Food"
// @Success 200 {object} menu.Status
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu/updateFood [put]
func (h *Handler) UpdateFood(c *gin.Context) {
	req := &pb.UpdateF{}
	if err := c.ShouldBindJSON(req); err != nil {
		h.Logger.Error("UpdateFood error: %v")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	resp, err := h.MenuClient.UpdateFood(context.Background(), req)
	if err != nil {
		h.Logger.Error("UpdateFood request error: %v")
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
// @Param request body menu.FoodId true "Delete Food"
// @Success 200 {object} menu.Status
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /menu/deleteFood/:id [delete]
func (h *Handler) DeleteFood(c *gin.Context) {
	req := &pb.FoodId{
		Id: c.Param("id"),
	}
	resp, err := h.MenuClient.DeleteFood(context.Background(), req)
	if err != nil {
		h.Logger.Error("DeleteFood request error: %v")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
