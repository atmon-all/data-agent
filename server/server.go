package server

import (
	"fmt"
	"net"

	"github.com/atmom/datagent/config"
	datagent "github.com/atmom/datagent/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// start server.
func Start(config *config.AgentConfig) {
	if config.Port < 1024 || config.Port > 65536 {
		logrus.Fatalf("Data agent server start error, invalid server port %d.", config.Port)
	}

	// listen tcp.
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.Fatalf("Data agent server start error %v.", err)
	}

	// create server.
	s := grpc.NewServer()
	datagent.RegisterDatagentServer(s, &service{Config: config})

	logrus.Infof("Data agent server start success, listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		logrus.Fatalf("Data agent server start error %v.", err)
	}
}
