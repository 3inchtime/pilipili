package main

import (
	"os"
	"os/signal"
	"pilipili/backend/internal/server"
	"pilipili/backend/internal/service"
	"syscall"
)

func main() {
	s := service.NewService()
	server.NewServer(s)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch <-c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		}
	}
}
