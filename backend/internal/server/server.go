package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"pilipili/backend/internal/service"
)

func NewServer(s *service.Service){
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