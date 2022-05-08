package server

import (
	"context"

	"github.com/atmom/datagent/config"
	datagent "github.com/atmom/datagent/grpc"
	"github.com/sirupsen/logrus"
)

type service struct {
	datagent.UnimplementedDatagentServer
	Config *config.AgentConfig
}

// Get data from data agent.
func (service *service) GetData(ctx context.Context, request *datagent.GetDataRequest) (*datagent.GetDataResponse, error) {
	logrus.Infof("Received: %v", request)
	return &datagent.GetDataResponse{Code: 0, Message: "success"}, nil
}
