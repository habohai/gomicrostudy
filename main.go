package main

import (
	"net/http"

	"github.com/micro/go-micro/web"
)

func main() {
	server := web.NewService(
		web.Address(":9060"),
	)

	server.HandleFunc("/", func(write http.ResponseWriter, request *http.Request) {
		write.Write([]byte("hello world"))
	})

	server.Run()
}
