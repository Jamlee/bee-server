module github.com/jamlee/bee-server

require (
	github.com/davyxu/cellnet v0.0.0-20181025100214-ebba2f27eee1
	github.com/davyxu/golog v0.0.0-20180706014138-e51e2504138b // indirect
	github.com/davyxu/goobjfmt v0.0.0-20180817064625-baf5de0715b1 // indirect
	github.com/hashicorp/consul v1.3.0
	github.com/hashicorp/go-cleanhttp v0.5.0 // indirect
	github.com/hashicorp/go-rootcerts v0.0.0-20160503143440-6bb64b370b90 // indirect
	github.com/hashicorp/serf v0.8.1 // indirect
	github.com/mitchellh/go-homedir v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/sirupsen/logrus v1.1.1
	github.com/urfave/cli v1.18.0
	go.etcd.io/etcd v3.3.10+incompatible // indirect
)

replace go.etcd.io/etcd v3.3.10+incompatible => github.com/etcd-io/etcd v3.3.10+incompatible
