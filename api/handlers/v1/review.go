package v1

import (
	pb "api_gateway/genproto/review"
	"api_gateway/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// @Summary Creates review
// @Description creates
// @Tags Reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param review body review.ReqCreateReview true "Review information"
// @Success 201 {object} review.ReviewInfo "Review informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /reviews/create [post]
func (h *HandlerV1) CreateReview(ctx *gin.Context) {

	req := pb.ReqCreateReview{}
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

	_, err = uuid.Parse(req.OrderId)
	if err != nil {
		h.log.Info("Invalid order id type uuid  ", zap.Any("input", req.KitchenId), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid order id type uuid  ",
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

	reviewInfo, err := h.reviewService.CreateReview(ctx, &req)
	if err != nil {
		h.log.Info("Error while creating review ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while creating review ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, reviewInfo)
}

// @Summary Get reviews
// @Description gets
// @Tags Reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "review id"
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Param review body review.ReqCreateReview true "Review information"
// @Success 201 {object} review.ReviewInfo "Review informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /reviews/{id} [get]
func (h *HandlerV1) GetReviewsByKitchenId(ctx *gin.Context) {

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
	req.Id = id

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

	reviews, err := h.reviewService.GetReviewsByKitchenId(ctx, &req)
	if err != nil {
		h.log.Error("Error while getting reviews by kitchenId ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while getting reviews by kitchenId ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, reviews)
}

// @Summary deletes review
// @Description deletes
// @Tags Reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "review id"
// @Param review body review.ReqCreateReview true "Review information"
// @Success 201 {object} review.ReviewInfo "Review informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /reviews/{id} [delete]
func (h *HandlerV1) DeleteReview(ctx *gin.Context) {

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

	reviews, err := h.reviewService.DeleteComment(ctx, &pb.Id{Id: id})
	if err != nil {
		h.log.Error("Error while getting reviews by kitchenId ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while getting reviews by kitchenId ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, reviews)
}
