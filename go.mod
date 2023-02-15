module github.com/containerd/containerd

go 1.18

require (
	github.com/containerd/ttrpc v1.1.1-0.20230127163717-32fab2374638
	github.com/containerd/typeurl/v2 v2.1.0
	github.com/google/go-cmp v0.5.9
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20230131230820-1c016267d619
	google.golang.org/grpc v1.52.3
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/prometheus/procfs v0.8.0 // indirect
	golang.org/x/net v0.5.0 // indirect
)

replace github.com/opencontainers/runtime-tools => github.com/opencontainers/runtime-tools v0.0.0-20221026201742-946c877fa809

replace github.com/AdaLogics/go-fuzz-headers => github.com/AdamKorcz/go-fuzz-headers-1 v0.0.0-20230111232327-1f10f66a31bf
