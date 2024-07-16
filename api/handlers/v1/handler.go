package v1

import (
	pbu "api_gateway/genproto/user"
	pbk "api_gateway/genproto/kitchen"
	"api_gateway/models"
	"api_gateway/pkg/connections"

	"go.uber.org/zap"
)

type HandlerV1 struct {
	log            *zap.Logger
	userService    pbu.UserServiceClient
	kitchenService pbk.KitchenClient
}

func NewHandlerV1(sysConfig *models.SystemConfig) *HandlerV1 {
	return &HandlerV1{
		log:         sysConfig.Logger,
		userService: connections.NewUserService(sysConfig),
		kitchenService: connections.NewKitchenService(sysConfig),
	}
}
