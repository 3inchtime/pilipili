package service

import (
	"context"
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
	panic("implement me")
}
