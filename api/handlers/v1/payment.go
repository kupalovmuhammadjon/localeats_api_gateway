package v1

import (
	pb "api_gateway/genproto/payment"
	"api_gateway/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Creates payment
// @Description creates
// @Tags Payments
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order body payment.ReqCreatePayment true "Payment information"
// @Success 201 {object} payment.PaymentInfo "Payment informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /payments/create [post]
func (h *HandlerV1) CreatePayment(ctx *gin.Context) {
	req := pb.ReqCreatePayment{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding ",
			Error:   err.Error(),
		})
		return
	}

	paymentInfo, err := h.paymentService.CreatePayment(ctx, &req)
	if err != nil {
		h.log.Info("Error while creating payment ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while creating payment ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, paymentInfo)
}
