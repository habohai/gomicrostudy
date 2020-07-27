package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func callAPI(addr string, path string, method string) (string, error) {
	req, err := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

func main() {
	reg := consul.NewRegistry(
		registry.Addrs("192.168.31.82:8500"),
	)

	getService, err := reg.GetService("prodservice")
	if err != nil {
		log.Fatal(err)
	}
	next := selector.RoundRobin(getService)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(node.Id, node.Address, node.Metadata)

	callRes, err := callAPI(node.Address, "/v1/prods", "GET")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(callRes)
}
