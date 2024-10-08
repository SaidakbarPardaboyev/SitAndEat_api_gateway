package handler

import (
	pb "api_gateway/genproto/users"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/swaggo/swag"

	"github.com/gin-gonic/gin"
)

// @Summary Get Profile
// @Description Get Profile
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "get Profile"
// @Success 200 {object} users.GetUser
// @Failure 400 {object} models.Error
// @Router /users/getProfile/{id} [get]
func (h *Handler) GetProfileById(c *gin.Context) {
	userId := c.Param("id")
	fmt.Println(userId)
	if _, err := uuid.Parse(userId); err != nil {
		h.Logger.Error(fmt.Sprintf("Invalid uuid?: %v", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid uuid?: %v", err.Error()),
		})
		return
	}
	req := pb.UserId{
		UserId: userId,
	}

	res, err := h.UserClient.GetProfile(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetProfile request error: %v", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Profile
// @Description Update Profile
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body users.UpdateProf true "Update Profile"
// @Success 200 {object} users.Status
// @Failure 400 {object} models.Error
// @Router /users/updateProfile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	req := pb.UpdateProf{}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateProfile error: %v", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.UserClient.UpdateProfile(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateProfile request error: %v", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully!",
	})
}

// @Summary Delete Profile
// @Description Delete Profile
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Delete Profile"
// @Success 200 {object} users.Status
// @Failure 400 {object} models.Error
// @Router /users/deleteProfile/{id} [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	id := c.Param("id")

	req := pb.UserId{
		UserId: id,
	}

	if _, err := uuid.Parse(id); err != nil {
		h.Logger.Error(fmt.Sprintf("Invalid uuid?: %v", err.Error()))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid uuid?: %v", err.Error()),
		})
		return
	}

	_, err := h.UserClient.DeleteProfile(c, &req)
	if err != nil {
		h.Logger.Error("DeleteProfile request error: %v", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile deleted successfully!",
	})
}
