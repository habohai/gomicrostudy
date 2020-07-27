package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haibeihabo/gomicrostudy/helper"
	"github.com/haibeihabo/gomicrostudy/prodservice"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("192.168.31.82:12379"),
	)

	ginRouter := gin.Default()

	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(context *gin.Context) {
			pr := helper.ProdsRequest{}
			if err := context.Bind(&pr); err != nil {
				pr = helper.ProdsRequest{Size: 2}
			}

			context.JSON(
				200,
				gin.H{
					"data": prodservice.NewProdList(pr.Size),
				},
			)
		})
	}

	server := web.NewService(
		web.Name("prodservice"),
		// web.Address(":9061"),
		web.Handler(ginRouter),
		web.Registry(reg),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	server.Init()
	server.Run()
}
