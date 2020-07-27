#! /bin/bash

# use for local test

go run /Users/haibei/study/goproject/gomicrostudy/prod/prod_main.go --server_address :9061 &
go run /Users/haibei/study/goproject/gomicrostudy/prod/prod_main.go --server_address :9062 &
go run /Users/haibei/study/goproject/gomicrostudy/prod/prod_main.go --server_address :9063 &
