package connections

import (
	pbk "api_gateway/genproto/kitchen"
	pbu "api_gateway/genproto/user"

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
