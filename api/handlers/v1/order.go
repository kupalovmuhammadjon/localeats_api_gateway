package v1

import (
	pb "api_gateway/genproto/order"
	"api_gateway/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// @Summary Creates order
// @Description creates
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order body order.ReqCreateOrder true "Order information"
// @Success 201 {object} order.OrderInfo "Order informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /orders/create [post]
func (h *HandlerV1) CreateOrder(ctx *gin.Context) {
	req := pb.ReqCreateOrder{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding ",
			Error:   err.Error(),
		})
		return
	}

	_, err = uuid.Parse(req.KitchenId)
	if err != nil {
		h.log.Info("Invalid kitchen id type uuid  ", zap.Any("input", req.KitchenId), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid kitchen id type uuid  ",
			Error:   err.Error(),
		})
		return
	}

	_, err = uuid.Parse(req.UserId)
	if err != nil {
		h.log.Info("Invalid user id type uuid  ", zap.Any("input", req.KitchenId), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid user id type uuid  ",
			Error:   err.Error(),
		})
		return
	}

	dish, err := h.orderService.CreateOrder(ctx, &req)
	if err != nil {
		h.log.Error("Error while creating order ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while creating order ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, dish)
}

// @Summary updates order
// @Description updates
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order body order.Status true "Status information"
// @Success 200 {object} order.StatusRes "Order informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /orders/{id}/update [put]
func (h *HandlerV1) UpdateOrderStatus(ctx *gin.Context) {
	req := pb.Status{}

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid  ", zap.Any("input", id), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid  ",
			Error:   err.Error(),
		})
		return
	}

	err = json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding ",
			Error:   err.Error(),
		})
		return
	}
	req.Id = id

	dish, err := h.orderService.UpdateOrderStatus(ctx, &req)
	if err != nil {
		h.log.Error("Error while UpdateOrderStatus ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while UpdateOrderStatus ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, dish)
}

// @Summary gets order
// @Description gets
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "order id"
// @Success 200 {object} order.OrderInfo "Order informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /orders/{id} [get]
func (h *HandlerV1) GetOrderById(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid  ", zap.Any("input", id), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid  ",
			Error:   err.Error(),
		})
		return
	}

	dish, err := h.orderService.GetOrderById(ctx, &pb.Id{Id: id})
	if err != nil {
		h.log.Error("Error while GetOrderById ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while GetOrderById ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, dish)
}

// @Summary gets order
// @Description gets
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "order id"
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {object} order.OrderInfo "Order informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /orders/user/{id} [get]
func (h *HandlerV1) GetOrdersForUser(ctx *gin.Context) {

	req := pb.Filter{}
	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid  ", zap.Any("input", id), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid  ",
			Error:   err.Error(),
		})
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil || limit <= 0 {
		h.log.Error("Invalid limit", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid limit parameter",
			Error:   err.Error(),
		})
		return
	}
	req.Limit = int32(limit)

	p, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || p <= 0 {
		h.log.Error("Invalid offset", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid offset parameter",
			Error:   err.Error(),
		})
		return
	}
	req.Page = int32(p)

	dish, err := h.orderService.GetOrdersForUser(ctx, &req)
	if err != nil {
		h.log.Error("Error while GetOrdersForUser ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while GetOrdersForUser ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, dish)
}

// @Summary gets order
// @Description gets
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "order id"
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {object} order.OrderInfo "Order informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /orders/chef/{id} [get]
func (h *HandlerV1) GetOrdersForChef(ctx *gin.Context) {

	req := pb.Filter{}
	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid  ", zap.Any("input", id), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid  ",
			Error:   err.Error(),
		})
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil || limit <= 0 {
		h.log.Error("Invalid limit", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid limit parameter",
			Error:   err.Error(),
		})
		return
	}
	req.Limit = int32(limit)

	p, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || p <= 0 {
		h.log.Error("Invalid offset", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid offset parameter",
			Error:   err.Error(),
		})
		return
	}
	req.Page = int32(p)

	dish, err := h.orderService.GetOrdersForChef(ctx, &req)
	if err != nil {
		h.log.Error("Error while GetOrdersForChef ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while GetOrdersForChef ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, dish)
}

// @Summary deletes order
// @Description deletes order
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} dish.Dishes "Dish informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /orders/{kitchen_id} [delete]
func (h *HandlerV1) DeleteOrder(ctx *gin.Context) {

	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid  ", zap.Any("input", id), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid  ",
			Error:   err.Error(),
		})
		return
	}

	dish, err := h.orderService.DeleteOrder(ctx, &pb.Id{Id: id})

	if err != nil {
		h.log.Error("Error while deleting order ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while deleting order ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, dish)
}
