module github.com/makegalxy/galxy/cmd/gateway

go 1.23.3

replace (
	github.com/makegalxy/galxy/pkg/services/gateway => ./../../pkg/services/gateway
	github.com/makegalxy/galxy/pkg/proto => ./../../pkg/proto
)

require (
	github.com/makegalxy/galxy/pkg/services/gateway v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.68.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.35.2-20241127180247-a33202765966.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0 // indirect
	github.com/makegalxy/galxy/pkg/proto v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)
