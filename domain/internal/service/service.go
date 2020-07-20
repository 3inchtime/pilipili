package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "pilipili/proto"
)

type Service struct {
	Router *gin.Engine
	RPC    pb.PilipiliClient
}

func NewService() *Service {
	s := &Service{
		Router: gin.Default(),
		RPC:    NewGRPC(),
	}
	return s
}

func NewGRPC() pb.PilipiliClient {
	Address := "127.0.0.1:23332"
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatal(err)
		logrus.Errorf("Start GRPC Error: %s", err.Error())
	}

	c := pb.NewPilipiliClient(conn)
	return c
}
