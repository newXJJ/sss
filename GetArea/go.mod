module sss/GetArea

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.

require (
	github.com/astaxie/beego v1.12.1
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/registry/consul/v2 v2.8.0
	github.com/micro/micro/v2 v2.9.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	sss/GetArea/handler v0.0.0
	sss/GetArea/proto/GetArea v0.0.0

)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace (
	sss/GetArea/handler => ./handler
	sss/GetArea/proto/GetArea => ./proto/GetArea
)
