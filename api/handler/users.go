package handler

import (
	"api_gateway/genproto/users"
	user "api_gateway/genproto/users"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary User profile ni ko'rsatish
// @Description Profile malumotlarini olish
// @ID User
// @Produce json
// @Success 200 {object} user.GetUser "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router //getProfile/:id [get]
func (h *Handler) GetProfile(c *gin.Context) {
	req := user.UserId{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading from body",
		})
		return
	}

	resp, err := h.UserClient.GetProfile(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	req := users.UpdateProf{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully!",
	})
}

func (h *Handler) DeleteProfile(c *gin.Context) {

}