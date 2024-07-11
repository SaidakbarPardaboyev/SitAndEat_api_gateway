package handler

import (
	"api_gateway/genproto/users"
	"net/http"

	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"

	"github.com/gin-gonic/gin"
)

// @Summary Get Profile
// @Description Get Profile
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} users.User
// @Failure 400 {object} models.Error
// @Router /users/getProfile/{id} [get]
func (h *Handler) GetProfileById(c *gin.Context) {
	claims, exists := c.Get("claims")

	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Claims not found",
		})
		return
	}
	id := claims.(jwt.MapClaims)["user_id"].(string)

	req := &users.UserId{
		UserId: id,
	}

	res, err := h.UserClient.GetProfile(c, req)
	if err != nil {
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
// @Param user body users.UpdateProf true "Update Profile"
// @Success 200 {object} users.UpdateProf
// @Failure 400 {object} models.Error
// @Router /users/updateProfile [put]

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

// @Summary Delete Profile
// @Description Delete Profile
// @Tags users
// @Accept json
// @Produce json
// @Param user body users.UserId true "Delete Profile"
// @Success 200 {object} users.UserId
// @Failure 400 {object} models.Error
// @Router /users/deleteProfile/{id} [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {

	req := users.UserId{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile deleted successfully!",
	})

}
