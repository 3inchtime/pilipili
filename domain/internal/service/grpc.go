package service

import (
	"github.com/gin-gonic/gin"
)

func(s *Service) GRPCGetVideoInfo (ctx *gin.Context){
	id := ctx.Query("id")

}
