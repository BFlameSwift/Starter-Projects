package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	v1 "hertz-demo/api/v1"
	"hertz-demo/middleware"
)

func InitRouter(h server.Hertz) error {
	h.Use(middleware.Cors())
	//TODO jwt authentication
	ApiGroup := h.Group("api")
	{
		ApiGroup.GET("/ping", v1.Hello)
	}
	return nil
}
