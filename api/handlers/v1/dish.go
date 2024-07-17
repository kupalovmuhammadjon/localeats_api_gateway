package v1

import (
	pb "api_gateway/genproto/dish"
	"api_gateway/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// @Summary Creates dish of kitchen
// @Description creates dish
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param dish body dish.ReqCreateDish true "Kitchen information"
// @Success 201 {object} dish.DishInfo "Dish informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /dishes/create [post]
func (h *HandlerV1) CreateDish(ctx *gin.Context) {
	req := pb.ReqCreateDish{}

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
		h.log.Info("Invalid id type uuid  ", zap.Any("input", req.KitchenId), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid  ",
			Error:   err.Error(),
		})
		return
	}

	dish, err := h.dishService.CreateDish(ctx, &req)
	if err != nil {
		h.log.Error("Error while creating dish ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while creating dish ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(201, dish)
}

// @Summary Update dish
// @Description Update dish of kitchen
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "dish id"
// @Param kitchen body dish.ReqUpdateDish true "Dish information"
// @Success 200 {object} dish.DishInfo "Dish informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /dishes/{id}/update [put]
func (h *HandlerV1) UpdateDish(ctx *gin.Context) {
	req := pb.ReqUpdateDish{}

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

	dish, err := h.dishService.UpdateDish(ctx, &req)

	if err != nil {
		h.log.Error("Error while updating dish ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while updating dish ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, dish)
}


// @Summary UpdateDishNutritionInfo
// @Description  UpdateDishNutritionInfo 
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "dish id"
// @Param kitchen body dish.NutritionInfo true "Dish information"
// @Success 200 {object} dish.DishInfo "Dish informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /dishes/{id}/updatenutritions [put]
func (h *HandlerV1) UpdateDishNutritionInfo(ctx *gin.Context) {
	req := pb.NutritionInfo{}

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

	dish, err := h.dishService.UpdateNutritionInfo(ctx, &req)

	if err != nil {
		h.log.Error("Error while updating dish NutritionInfo", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while updating dish NutritionInfo",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, dish)
}

// @Summary gets dish by id
// @Description gets dish by id
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "dish id"
// @Success 200 {object} dish.DishInfo "Dish informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /dishes/{id} [get]
func (h *HandlerV1) GetDishById(ctx *gin.Context) {

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

	dish, err := h.dishService.GetDishById(ctx, &pb.Id{Id: id})

	if err != nil {
		h.log.Error("Error while getting dish by id ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while getting dish by id ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, dish)
}

// @Summary gets dishes
// @Description gets dishes of the kitchen
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "kitchen id"
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {object} dish.Dishes "Dish informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /dishes/all/{id} [get]
func (h *HandlerV1) GetDishesByKitchenId(ctx *gin.Context) {

	req := pb.Pagination{}

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

	dish, err := h.dishService.GetDishes(ctx, &req)

	if err != nil {
		h.log.Error("Error while getting dishes by kitchen_id ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while getting dishes by kitchen_id ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, dish)
}

// @Summary deletes dish
// @Description deletes dish
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Success 200 {object} dish.Dishes "Dish informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /dishes/{kitchen_id} [delete]
func (h *HandlerV1) DeleteDish(ctx *gin.Context) {

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

	dish, err := h.dishService.DeleteDish(ctx, &pb.Id{Id: id})

	if err != nil {
		h.log.Error("Error while deleting dish ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while deleting dish ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, dish)
}

// @Summary gets dishes
// @Description gets dishes of the kitchen
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id path string true "user id"
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {object} dish.Recommendations "Dish informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /dishes/recommendations/{user_id} [get]
func (h *HandlerV1) RecommendDishes(ctx *gin.Context) {

	req := pb.Filter{}

	id := ctx.Param("user_id")
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

	dishes, err := h.dishService.RecommendDishes(ctx, &req)

	if err != nil {
		h.log.Error("Error while getting recommended dishes by kitchen_id ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while getting recommended dishes by kitchen_id ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, dishes)
}