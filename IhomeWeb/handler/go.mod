module sss/IhomeWeb/handler

go 1.14

require (
	path/to/service/proto/IhomeWeb v0.0.0
	path/to/service/proto/GetArea v0.0.0
)
replace (
    path/to/service/proto/IhomeWeb => ../../../www/example/com/pb/proto/user
    path/to/service/proto/GetArea => ../../GetArea/proto/GetArea
)
