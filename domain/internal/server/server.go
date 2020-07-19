package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"pilipili/domain/internal/service"
)

func NewServer(s *service.Service){
	s.InitRoutes()
	server := http.Server{
		Addr: ":22222",
		Handler: s.Router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Errorf("HTTP Server Error: %s", err.Error())
		}
	}()
}

