module sss/GetImageCd

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/astaxie/beego v1.12.2
	github.com/dev-submail/submail_go_sdk v0.0.0-20190807202824-e1fadb4a3bac // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/registry/consul/v2 v2.8.0
	github.com/micro/micro/v2 v2.9.2 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	sss/GetImageCd/handler v0.0.0
	sss/GetImageCd/proto/GetImageCd v0.0.0
)

replace (
	sss/GetImageCd/handler => ./handler
	sss/GetImageCd/proto/GetImageCd => ./proto/GetImageCd
)
