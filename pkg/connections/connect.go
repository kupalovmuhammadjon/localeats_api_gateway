package connections

import (
	pbk "api_gateway/genproto/kitchen"
	pbu "api_gateway/genproto/user"
	pbd "api_gateway/genproto/dish"
	pbo "api_gateway/genproto/order"
	pbp "api_gateway/genproto/payment"
	pbr "api_gateway/genproto/review"

	"api_gateway/models"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserService(sysConfig *models.SystemConfig) pbu.UserServiceClient {
	conn, err := grpc.NewClient(sysConfig.Config.AUTH_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sysConfig.Logger.Fatal("Failed to connect auth service user client ", zap.Error(err))
		return nil
	}

	return pbu.NewUserServiceClient(conn)
}

func NewKitchenService(sysConfig *models.SystemConfig) pbk.KitchenClient {
	conn, err := grpc.NewClient(sysConfig.Config.AUTH_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sysConfig.Logger.Fatal("Failed to connect auth service kitchen client ", zap.Error(err))
		return nil
	}

	return pbk.NewKitchenClient(conn)
}

func NewDishService(sysConfig *models.SystemConfig) pbd.DishClient {
	conn, err := grpc.NewClient(sysConfig.Config.ORDER_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sysConfig.Logger.Fatal("Failed to connect order service dish client ", zap.Error(err))
		return nil
	}

	return pbd.NewDishClient(conn)
}

func NewOrderService(sysConfig *models.SystemConfig) pbo.OrderClient {
	conn, err := grpc.NewClient(sysConfig.Config.ORDER_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sysConfig.Logger.Fatal("Failed to connect order service order client ", zap.Error(err))
		return nil
	}

	return pbo.NewOrderClient(conn)
}

func NewPaymentService(sysConfig *models.SystemConfig) pbp.PaymentClient {
	conn, err := grpc.NewClient(sysConfig.Config.ORDER_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sysConfig.Logger.Fatal("Failed to connect order service payment client ", zap.Error(err))
		return nil
	}

	return pbp.NewPaymentClient(conn)
}

func NewReviewService(sysConfig *models.SystemConfig) pbr.ReviewClient {
	conn, err := grpc.NewClient(sysConfig.Config.ORDER_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sysConfig.Logger.Fatal("Failed to connect order service review client ", zap.Error(err))
		return nil
	}

	return pbr.NewReviewClient(conn)
}