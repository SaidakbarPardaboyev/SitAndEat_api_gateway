package handler

import (
	"api_gateway/genproto/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(c *gin.Context) {

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
