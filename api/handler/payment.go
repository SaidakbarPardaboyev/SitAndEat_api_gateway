package handler

import (
	pb "api_gateway/genproto/payment"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a payment
// @Description Create a new payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param request body payment.CreatePayment true "Create Payment Request"
// @Success 200 {object} payment.Status
// @Failure 400 {object} models.Error
// @Router /payments/createPayment [post]
func (h *Handler) CreatePayments(c *gin.Context) {
	req := pb.CreatePayment{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Create payment error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	resp,err:=h.PaymentClient.CreatePayments(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		h.Logger.Error("Create payment request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get payment status by ID
// @Description Get payment status by ID
// @Tags payments
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Success 200 {object} payment.GetByIdResponse
// @Failure 400 {object} models.Error
// @Router /payments/getPaymentStatus/{id} [get]
func (h *Handler) GetPaymentStatusById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		h.Logger.Error("ID parameter is missing")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID parameter is missing",
		})
		return
	}

	req := pb.GetById{
		Id: id,
	}

	resp, err := h.PaymentClient.GetPaymentStatusById(context.Background(), &req)
	if err != nil {
		h.Logger.Error("GetPaymentStatusById request error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Update a payment
// @Description Update an existing payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param request body payment.UpdatePayment true "Update Payment Request"
// @Success 200 {object} payment.Status
// @Failure 400 {object} models.Error
// @Router /payments/updatePayment [put]
func (h *Handler) UpdatePayments(c *gin.Context) {
    req := pb.UpdatePayment{}
    if err := c.ShouldBindJSON(&req); err != nil {
        h.Logger.Error("UpdatePayments error: ", err)
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    resp, err := h.PaymentClient.UpdatePayments(context.Background(), &req)
    if err != nil {
        h.Logger.Error("UpdatePayments request error: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, resp)
}
