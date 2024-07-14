package connections

import (
	pbu "api_gateway/genproto/user"
	"api_gateway/models"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewUserService(sysConfig *models.SystemConfig) pbu.UserServiceClient {
	conn, err := grpc.NewClient(sysConfig.Config.AUTH_SERVICE_PORT)
	if err != nil {
		sysConfig.Logger.Fatal("Failed to connect auth servoce", zap.Error(err))
		return nil
	}

	u := pbu.NewUserServiceClient(conn)

	return u
}
