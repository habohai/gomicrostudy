package main

import (
	"github.com/gin-gonic/gin"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	reg := consul.NewRegistry(
		registry.Addrs("192.168.31.82:8500"),
	)

	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/user", func(context *gin.Context) {
		context.JSON(200, "user api")
	})

	ginRouter.Handle("GET", "/news", func(context *gin.Context) {
		context.JSON(200, "news api")
	})

	server := web.NewService(
		web.Name("prodservice"),
		web.Address(":9060"),
		web.Handler(ginRouter),
		web.Registry(reg),
	)

	server.Run()
}
