package service

import (
	"github.com/gin-gonic/gin"
	"pilipili/backend/dao"
)

type Service struct {
	dao *dao.Dao
	Router *gin.Engine
}

func NewService() *Service {
	s := &Service{
		dao: dao.NewDao(),
		Router : gin.Default(),
	}
	s.initRoutes()
	return s
}