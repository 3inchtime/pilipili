package service

import (
	"context"
	"pilipili/backend/model"
	pb "pilipili/proto"
)

func (s *Service) RpcGetVideoInfo(ctx context.Context, request *pb.GetVideoRequest) (*pb.VideoInfo, error) {
	id := request.Id
	v := s.dao.GetVideoInfo(id)

	resp := new(pb.VideoInfo)

	resp.Id = id
	resp.Title = v.Title
	resp.Note = v.Note
	resp.Pic = v.PicPath
	resp.Video = v.VideoPath

	return resp, nil
}

func (s *Service) RpcCreateVideoInfo(ctx context.Context, info *pb.VideoInfo) (*pb.CreateVideoReplay, error) {
	title := info.Title
	note := info.Note
	pic := info.Pic
	video := info.Video

	v := new(model.Video)
	v.Title = title
	v.Note = note
	v.PicPath = pic
	v.VideoPath = video

	resp := new(pb.CreateVideoReplay)
	resp.Status = "OK"
	return resp, nil
}
