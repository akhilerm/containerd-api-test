module github.com/containerd/containerd

go 1.18

require (
	github.com/containerd/ttrpc v1.1.1-0.20220420014843-944ef4a40df3
	github.com/containerd/typeurl v1.0.3-0.20220422153119-7f6e6d160d67
	github.com/google/go-cmp v0.5.9
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/sys v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20221206210731-b1a01be3a5f6
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/text v0.5.0 // indirect
)

replace github.com/opencontainers/runtime-tools => github.com/opencontainers/runtime-tools v0.0.0-20221026201742-946c877fa809
