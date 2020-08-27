module sss/IhomeWeb

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require github.com/micro/go-micro/v2 v2.9.0

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868 // indirect
	github.com/astaxie/beego v1.12.1
	github.com/dev-submail/submail_go_sdk v0.0.0-20190807202824-e1fadb4a3bac // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/julienschmidt/httprouter v1.2.0
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-plugins/registry/consul/v2 v2.8.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	path/to/service/proto/GetImageCd v0.0.0 // indirect
	sss/IhomeWeb/handler v0.0.0
)

replace (
	path/to/service/proto/GetArea => ../GetArea/proto/GetArea
	path/to/service/proto/GetImageCd => ../GetImageCd/proto/GetImageCd
	path/to/service/proto/IhomeWeb => ../../www/example/com/pb/proto/user
	sss/IhomeWeb/handler => ./handler
)
