package connections

import (
	pbu "api_gateway/genproto/user"
	"api_gateway/models"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserService(sysConfig *models.SystemConfig) pbu.UserServiceClient {
	conn, err := grpc.NewClient(sysConfig.Config.AUTH_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sysConfig.Logger.Fatal("Failed to connect auth service", zap.Error(err))
		return nil
	}

	u := pbu.NewUserServiceClient(conn)

	return u
}
