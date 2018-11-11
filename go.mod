module github.com/jamlee/bee-server

require (
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973 // indirect
	github.com/coreos/bbolt v1.3.0 // indirect
	github.com/coreos/etcd v3.3.10+incompatible // indirect
	github.com/coreos/go-semver v0.2.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20181031085051-9002847aa142 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/davyxu/cellnet v0.0.0-20181109151218-c9ddeeedf7ba
	github.com/davyxu/golog v0.0.0-20180706014138-e51e2504138b // indirect
	github.com/davyxu/goobjfmt v0.0.0-20180817064625-baf5de0715b1 // indirect
	github.com/davyxu/protoplus v0.0.0-20181101064152-7f5ef06bcaec // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gogo/protobuf v1.1.1 // indirect
	github.com/google/btree v0.0.0-20180813153112-4030bb1f1f0c // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.5.1 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_golang v0.9.1 // indirect
	github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910 // indirect
	github.com/prometheus/common v0.0.0-20181109100915-0b1957f9d949 // indirect
	github.com/prometheus/procfs v0.0.0-20181005140218-185b4288413d // indirect
	github.com/sirupsen/logrus v1.1.1
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20171017195756-830351dc03c6 // indirect
	github.com/ugorji/go/codec v0.0.0-20181022190402-e5e69e061d4f // indirect
	github.com/urfave/cli v1.18.0
	github.com/xiang90/probing v0.0.0-20160813154853-07dd2e8dfe18 // indirect
	go.etcd.io/etcd v3.3.10+incompatible
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1 // indirect
	golang.org/x/net v0.0.0-20181108082009-03003ca0c849 // indirect
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c // indirect
	google.golang.org/grpc v1.16.0 // indirect
	gopkg.in/yaml.v2 v2.2.1 // indirect
)

replace (
	github.com/coreos/bbolt v1.3.0 => github.com/coreos/bbolt v1.3.1-etcd.8
	go.etcd.io/etcd v3.3.10+incompatible => github.com/etcd-io/etcd v3.3.10+incompatible
)
