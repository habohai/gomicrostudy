#! /bin/bash

# use for protoc to go

cd models/protos
protoc --micro_out=../ --go_out=../ prods.proto
cd ../../