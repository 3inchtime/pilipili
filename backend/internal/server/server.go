package server

import (
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"pilipili/backend/internal/service"
	pb "pilipili/proto"
)

func NewHttpServer(s *service.Service) {
	s.InitRoutes()
	server := http.Server{
		Addr: ":23333",
		Handler: s.Router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Errorf("HTTP Server Error: %s", err.Error())
		}
	}()
}

func NewGRPCServer(s *service.Service) {
	Address := "127.0.0.1:23332"
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		logrus.Errorf("Listen Error: %s", err.Error())
		panic(err)
	}
	pb.RegisterPilipiliServer(s.RPC, s)
	s.RPC.Serve(listen)
}