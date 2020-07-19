package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	Router *gin.Engine
	HttpClient *http.Client
}

func NewService() *Service {
	s := &Service{
		Router : gin.Default(),
		HttpClient: &http.Client{},
	}
	s.initRoutes()
	return s
}