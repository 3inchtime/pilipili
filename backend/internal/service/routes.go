package service

func(s *Service) initRoutes() {
	s.Router.GET("/video", s.GetVideoInfo)
	s.Router.POST("/video", s.CreateVideo)
}