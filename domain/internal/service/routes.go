package service

func(s *Service) initRoutes() {
	httpRoutes := s.Router.Group("/http")
	{
		httpRoutes.GET("/get_video", s.HttpGetVideoInfo)
		httpRoutes.POST("/create_video", s.HttpCreateVideo)
	}
}
