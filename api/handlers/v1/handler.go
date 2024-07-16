package v1

import (
	pbd "api_gateway/genproto/dish"
	pbk "api_gateway/genproto/kitchen"
	pbo "api_gateway/genproto/order"
	pbp "api_gateway/genproto/payment"
	pbr "api_gateway/genproto/review"
	pbu "api_gateway/genproto/user"
	"api_gateway/models"
	"api_gateway/pkg/connections"

	"go.uber.org/zap"
)

type HandlerV1 struct {
	log            *zap.Logger
	userService    pbu.UserServiceClient
	kitchenService pbk.KitchenClient
	dishService    pbd.DishClient
	orderService   pbo.OrderClient
	paymentService pbp.PaymentClient
	reviewService  pbr.ReviewClient
}

func NewHandlerV1(sysConfig *models.SystemConfig) *HandlerV1 {
	return &HandlerV1{
		log:            sysConfig.Logger,
		userService:    connections.NewUserService(sysConfig),
		kitchenService: connections.NewKitchenService(sysConfig),
		dishService:    connections.NewDishService(sysConfig),
		orderService:   connections.NewOrderService(sysConfig),
		paymentService: connections.NewPaymentService(sysConfig),
		reviewService:  connections.NewReviewService(sysConfig),
	}
}
