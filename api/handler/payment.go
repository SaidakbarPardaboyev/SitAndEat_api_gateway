package handler

import (
	pb "api_gateway/genproto/payment"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a payment
// @Description Create a new payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param request body pb.CreatePayment true "Create Payment Request"
// @Success 200 {object} pb.Status
// @Failure 400 {object} gin.H{"error": "string"}
// @Router /payments [post]
func (h *Handler) CreatePayments(c *gin.Context) {
	req := pb.CreatePayment{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		log.Fatalf("Create payment error: %v", err)
		return
	}

	resp,err:=h.PaymentClient.CreatePayments(context.Background(), &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		log.Fatalf("Create payment request error: %v", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get payment status by ID
// @Description Get payment status by ID
// @Tags payments
// @Accept  json
// @Produce  json
// @Param request body pb.GetById true "Get Payment By ID Request"
// @Success 200 {object} pb.GetByIdResponse
// @Failure 400 {object} gin.H{"error": "string"}
// @Router /payments/status [get]
func (h *Handler) GetPaymentStatusById(c *gin.Context){
	req:=pb.GetById{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		log.Fatalf("GetbyId payment error: %v", err)
		return
	}

	resp,err:=h.PaymentClient.GetPaymentStatusById(context.Background(), &req)
	if err!=nil{
		c.JSON(500,gin.H{
			"error":err.Error(),
		})
	}

	c.JSON(200,resp)
}

// @Summary Update a payment
// @Description Update an existing payment
// @Tags payments
// @Accept  json
// @Produce  json
// @Param request body pb.UpdatePayment true "Update Payment Request"
// @Success 200 {object} pb.Status
// @Failure 400 {object} gin.H{"error": "string"}
// @Router /payments [put]
func (h *Handler) UpdatePayments(c *gin.Context){
	req:=pb.UpdatePayment{}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
	}

	resp,err:=h.PaymentClient.UpdatePayments(context.Background(), &req)
	if err!=nil{
		c.JSON(500,gin.H{
			"error":err.Error(),
		})
	}
	c.JSON(200,resp)
}