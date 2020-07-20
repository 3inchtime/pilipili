package service

func (s *Service) InitRoutes() {
	httpRoutes := s.Router.Group("/http")
	{
		httpRoutes.GET("/get_video", s.HttpGetVideoInfo)
		httpRoutes.POST("/create_video", s.HttpCreateVideo)
	}
	grpcRoutes := s.Router.Group("/grpc")
	{
		grpcRoutes.GET("/get_video", s.GRPCGetVideoInfo)
		grpcRoutes.POST("/create_video", s.GRPCCreateVideo)
	}
}
