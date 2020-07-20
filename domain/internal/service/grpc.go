package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	pb "pilipili/proto"
)

func (s *Service) GRPCGetVideoInfo(ctx *gin.Context) {
	rpcRequest := new(pb.GetVideoRequest)
	rpcRequest.Id = ctx.Query("id")
	v, err := s.RPC.RpcGetVideoInfo(context.Background(), rpcRequest)
	if err != nil {
		logrus.Errorf("RPC Server Error: %s", err.Error())
	}
	if v != nil {
		ctx.JSON(http.StatusOK, gin.H{"id": v.Id, "title": v.Title, "note": v.Note, "pic": v.Pic, "video": v.Video})
	} else {
		ctx.String(http.StatusInternalServerError, "Failed")
	}
}


func (s *Service) GRPCCreateVideo(ctx *gin.Context) {
	rpcRequest := new(pb.VideoInfo)
	rpcRequest.Title = ctx.PostForm("title")
	rpcRequest.Note = ctx.PostForm("note")
	rpcRequest.Pic = ctx.PostForm("pic")
	rpcRequest.Video = ctx.PostForm("video")

	v, err := s.RPC.RpcCreateVideoInfo(context.Background(), rpcRequest)
	if err != nil {
		logrus.Errorf("RPC Server Error: %s", err.Error())
	}
	if v != nil {
		ctx.String(http.StatusOK, v.Status)
	} else {
		ctx.String(http.StatusInternalServerError, "Failed")
	}
}
