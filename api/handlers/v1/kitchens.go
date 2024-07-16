package v1

import (
	pb "api_gateway/genproto/kitchen"
	"api_gateway/models"
	"api_gateway/pkg/validations"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// @Summary Creates kitchen of user
// @Description creates kitchen
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param kitchen body kitchen.ReqCreateKitchen true "Kitchen information"
// @Success 200 {object} user.User "Profile informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /kitchens/create [post]
func (h *HandlerV1) CreateKitchen(ctx *gin.Context) {
	req := pb.ReqCreateKitchen{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding ",
			Error: err.Error(),
		})
		return
	}

	_, err = uuid.Parse(req.OwnerId)
	if err != nil {
		h.log.Info("Invalid owner id type uuid ", zap.Any("input", req.OwnerId), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid owner id type uuid   ",
			Error: err.Error(),
		})
		return
	}

	err = validations.ValidatePhoneNumber(req.PhoneNumber)
	if err != nil {
		h.log.Info("invalid phone number ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "invalid phone number ",
			Error: err.Error(),
		})
		return
	}
	kitchen, err := h.kitchenService.CreateKitchen(ctx, &req)

	if err != nil {
		h.log.Error("Error while creating kitchen ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while creating kitchen ",
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(201, kitchen)
}


// @Summary Creates kitchen of user
// @Description creates kitchen 
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Kitchen id"
// @Param kitchen body kitchen.ReqCreateKitchen true "Kitchen information"
// @Success 200 {object} user.User "Profile informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /kitchens/{id}/update [put]
func (h *HandlerV1) UpdateKitchen(ctx *gin.Context) {
	req := pb.ReqUpdateKitchen{}

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid  ", zap.Any("input", id), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid uuid  ",
			Error: err.Error(),
		})
		return
	}

	err = json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding ",
			Error: err.Error(),
		})
		return
	}

	req.Id = id

	kitchen, err := h.kitchenService.UpdateKitchen(ctx, &req)

	if err != nil {
		h.log.Error("Error while updating kitchen ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while updating kitchen ",
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(200, kitchen)
}

// @Summary Creates kitchen of user
// @Description creates kitchen 
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Kitchen id"
// @Success 200 {object} kitchen.KitchenInfo "Profile informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /kitchens/{id} [get]
func (h *HandlerV1) GetKitchenById(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid uuid ", zap.Any("input", id), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid uuid  ",
			Error: err.Error(),
		})
		return
	}

	kitchen, err := h.kitchenService.GetKitchenById(ctx, &pb.Id{Id: id})

	if err != nil {
		h.log.Error("Error while getting kitchen  by id ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while getting kitchen  by id",
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(200, kitchen)
}

// @Summary Get list of kitchens
// @Description Retrieves a paginated list of kitchens
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {object} kitchen.Kitchens "List of kitchens"
// @Failure 401 {object} models.Error "No Auth"
// @Failure 400 {object} models.Error "Invalid inputs"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /kitchens [get]
func (h *HandlerV1) GetKitchens(ctx *gin.Context) {
	req := pb.Pagination{}

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

	kitchens, err := h.kitchenService.GetKitchens(ctx, &req)
	if err != nil {
		h.log.Error("Error while getting kitchens", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while getting kitchens",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, kitchens)
}

// @Summary Get list of kitchens
// @Description Retrieves a paginated list of kitchens
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param name query string false "Name"
// @Param cuisine_type query string false "Cuisine Type"
// @Param rating query number true "Rating"
// @Param address query string false "Address"
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {object} kitchen.Kitchens "List of kitchens"
// @Failure 401 {object} models.Error "No Auth"
// @Failure 400 {object} models.Error "Invalid inputs"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /kitchens/search [get]
func (h *HandlerV1) SearchKitchens(ctx *gin.Context) {
	req := pb.Search{}

	req.Name = ctx.Query("name")
	req.CuisineType = ctx.Query("cuisine_type")

	rating, err := strconv.ParseFloat(ctx.Query("rating"), 32)
	if err == nil {
		req.Rating = float32(rating)
	}

	req.Address = ctx.Query("address")

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

	kitchens, err := h.kitchenService.SearchKitchens(ctx, &req)
	if err != nil {
		h.log.Error("Error while getting kitchens", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while getting kitchens",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, kitchens)
}

// @Summary Get list of kitchens
// @Description Retrieves a paginated list of kitchens
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Id"
// @Success 200 {object} kitchen.Kitchens "List of kitchens"
// @Failure 401 {object} models.Error "No Auth"
// @Failure 400 {object} models.Error "Invalid inputs"
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /kitchens/{id}/delete [delete]
func (h *HandlerV1) DeleteKitchen(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid uuid ", zap.Any("input", id), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid uuid  ",
			Error: err.Error(),
		})
		return
	}

	kitchens, err := h.kitchenService.DeleteKitchen(ctx, &pb.Id{Id: id})
	if err != nil {
		h.log.Error("Error while deleting kitchens", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while deleting kitchens",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, kitchens)
}
