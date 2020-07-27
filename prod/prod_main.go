package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haibeihabo/gomicrostudy/prodservice"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	reg := consul.NewRegistry(
		registry.Addrs("192.168.31.82:8500"),
	)

	ginRouter := gin.Default()

	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("GET", "/prods", func(context *gin.Context) {
			context.JSON(200, prodservice.NewProdList(5))
		})
	}

	server := web.NewService(
		web.Name("prodservice"),
		// web.Address(":9061"),
		web.Handler(ginRouter),
		web.Registry(reg),
	)
	server.Init()
	server.Run()
}
