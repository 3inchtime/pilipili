package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pilipili/backend/model"
)

func(s *Service) HttpGetVideoInfo (ctx *gin.Context){
	id := ctx.Query("id")
	v := s.dao.GetVideoInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
		"title": v.Title,
		"note": v.Note,
		"pic": v.PicPath,
		"video": v.VideoPath,
	})
}

func (s *Service) HttpCreateVideo (ctx *gin.Context){
	title := ctx.PostForm("title")
	note := ctx.PostForm("note")
	pic := ctx.PostForm("pic")
	video := ctx.PostForm("video")

	v := new(model.Video)
	v.Title = title
	v.Note = note
	v.PicPath = pic
	v.VideoPath = video

	s.dao.CreateNewVideo(v)

	ctx.String(http.StatusOK, "OK")
}

