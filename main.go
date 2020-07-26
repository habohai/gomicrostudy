package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/user", func(context *gin.Context) {
		context.JSON(200, "user api")
	})

	ginRouter.Handle("GET", "/news", func(context *gin.Context) {
		context.JSON(200, "news api")
	})

	server := web.NewService(
		web.Address(":9060"),
		web.Handler(ginRouter),
	)

	server.Run()
}
