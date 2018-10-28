module github.com/jamlee/bee-server

require (
	github.com/coreos/etcd v3.3.10+incompatible // indirect
	github.com/davyxu/cellnet v0.0.0-20181025100214-ebba2f27eee1
	github.com/davyxu/golog v0.0.0-20180706014138-e51e2504138b // indirect
	github.com/davyxu/goobjfmt v0.0.0-20180817064625-baf5de0715b1 // indirect
	github.com/davyxu/protoplus v0.0.0-20181026032816-9e740529cb1e // indirect
	github.com/gogo/protobuf v1.1.1 // indirect
	github.com/mitchellh/go-homedir v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/sirupsen/logrus v1.1.1
	github.com/urfave/cli v1.18.0
	go.etcd.io/etcd v3.3.10+incompatible
	google.golang.org/grpc v1.16.0 // indirect
)

replace go.etcd.io/etcd v3.3.10+incompatible => github.com/etcd-io/etcd v3.3.10+incompatible
