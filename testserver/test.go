package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/haibeihabo/gomicrostudy/models"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	myhttp "github.com/micro/go-plugins/client/http"
)

func callAPI2(s selector.Selector) {
	myclient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)

	req := myclient.NewRequest("prodservice", "/v1/prods", models.ProdsRequest{Size: 3})

	var rsp models.ProdListResponse
	err := myclient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Println("-------")
		log.Fatal(err)
	}

	log.Println(rsp.GetData())
}

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
	reg := etcd.NewRegistry(
		registry.Addrs("192.168.31.82:12379"),
	)

	mySelector := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)

	callAPI2(mySelector)
}
