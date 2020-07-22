package service

func (s *Service) InitRoutes() {
	httpRoutes := s.Router.Group("/http")
	{
		httpRoutes.GET("/get_video", s.HttpGetVideoInfo)
		httpRoutes.POST("/create_video", s.HttpCreateVideo)
	}
	gRPCRoutes := s.Router.Group("/grpc")
	{
		gRPCRoutes.GET("/get_video", s.GRPCGetVideoInfo)
		gRPCRoutes.POST("/create_video", s.GRPCCreateVideo)
	}
}