package v1

import (
	pb "api_gateway/genproto/user"
	"api_gateway/models"
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
// @Param id path string true "User id"
// @Success 200 {object} user.User "Profile informations"
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
