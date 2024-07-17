package v1

import (
	pb "api_gateway/genproto/user"
	"api_gateway/models"
	"api_gateway/pkg/validations"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// @Summary Gets User Profile by id
// @Description get full profile information
// @Tags Users
// @ID getprofile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User id"
// @Success 200 {object} user.User "Profile informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /users/{id}/profile [get]
func (h *HandlerV1) GetProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid",
			Error:   err.Error(),
		})
		return
	}

	req := &pb.Id{Id: id}
	_, err = h.userService.ValidateUserId(ctx, req)
	if err != nil {
		h.log.Info("Invalid user id ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid User Id",
			Error:   err.Error(),
		})
		return
	}

	user, err := h.userService.GetProfile(ctx, req)
	if err != nil {
		h.log.Error("Invalid user id ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Invalid User Id",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

// @Summary Gets User Profile by id
// @Description get full profile information
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User id"
// @Param user body user.ReqUpdateUser true "User id"
// @Success 200 {object} user.User "Profile informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /users/{id}/update [put]
func (h *HandlerV1) UpdateProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid",
			Error:   err.Error(),
		})
		return
	}

	req := &pb.ReqUpdateUser{}

	err = json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding ",
			Error:   err.Error(),
		})
		return
	}

	err = validations.ValidateEmail(req.Email)
	if err != nil {
		h.log.Info("Error in validating email ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid email",
			Error:   err.Error(),
		})
		return
	}
	err = validations.ValidatePhoneNumber(req.PhoneNumber)
	if err != nil {
		h.log.Info("Error in validating phone number ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid phone number",
			Error:   err.Error(),
		})
		return
	}
	req.Id = id
	user, err := h.userService.UpdateProfile(ctx, req)
	if err != nil {
		h.log.Error("failed to update user profile ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "failed to update user profile ",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

// @Summary deletes
// @Description deletes
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User id"
// @Success 200 {object} user.Status "About deletion"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /users/{id}/delete [delete]
func (h *HandlerV1) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid",
			Error:   err.Error(),
		})
		return
	}

	user, err := h.userService.DeleteUser(ctx, &pb.Id{Id: id})
	if err != nil {
		h.log.Error("failed to delete user ", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "failed to delete user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

// @Summary Gets User Preferences by id
// @Description get Preferences information
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User id"
// @Param user body user.Preferences true "User id"
// @Success 200 {object} user.PreferencesRes "Profile informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /users/{id}/updatepreferences [put]
func (h *HandlerV1) UpdateUserPreferences(ctx *gin.Context){

	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid",
			Error:   err.Error(),
		})
		return
	}

	req := &pb.Preferences{}

	err = json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding",
			Error:   err.Error(),
		})
		return
	}
	req.UserId = id

	res, err := h.userService.UpdateUserPreferences(ctx, req)
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}

// @Summary Gets User Preferences by id
// @Description get Preferences information
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User id"
// @Success 200 {object} user.PreferencesRes "Profile informations"
// @Failure 401 {object} models.Error "No Auth thats the problem "
// @Failure 400 {object} models.Error "Invalid inputs can result to "
// @Failure 500 {object} models.Error "Something went wrong in server"
// @Router /users/{id}/preferences [get]
func (h *HandlerV1) GetUserPreference(ctx *gin.Context){

	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		h.log.Info("Invalid id type uuid", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid id type uuid",
			Error:   err.Error(),
		})
		return
	}

	res, err := h.userService.GetUserPreference(ctx, &pb.Id{Id: id})
	if err != nil {
		h.log.Info("Error while decoding ", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Error while decoding",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}