bee-server
========
[![CodeFactor](https://www.codefactor.io/repository/github/jamlee/bee-server/badge)](https://www.codefactor.io/repository/github/jamlee/bee-server)
[![Build Status](https://travis-ci.com/Jamlee/bee-server.svg?branch=master)](https://travis-ci.com/Jamlee/bee-server)

## Building

`make`

## Running

```
# 启动主服务器 etcd-1
./bin/bee-server --name etcd-1 \
  --peer-url=http://127.0.0.1:2381 \
  --advertise-client-url=http://127.0.0.1:2379 \
  --initial-cluster-token=etcd \
  --initial-cluster etcd-1=http://127.0.0.1:2380,etcd-2=http://127.0.0.1:2381,etcd-3=http://127.0.0.1:2382 \
  server --master-port=20001 --web-port=10001

# 启动主服务器 etcd-2
./bin/bee-server --name etcd-2 \
  --peer-url=http://127.0.0.1:2381 \
  --advertise-client-url=http://127.0.0.1:2378 \
  --initial-cluster-token=etcd \
  --initial-cluster etcd-1=http://127.0.0.1:2380,etcd-2=http://127.0.0.1:2381,etcd-3=http://127.0.0.1:2382 \
  server --master-port=20002 --web-port=10002

# 启动主服务器 etcd-3
./bin/bee-server --name etcd-3 \
  --peer-url=http://127.0.0.1:2382 \
  --advertise-client-url=http://127.0.0.1:2377 \
  --initial-cluster-token=etcd \
  --initial-cluster etcd-1=http://127.0.0.1:2380,etcd-2=http://127.0.0.1:2381,etcd-3=http://127.0.0.1:2382 \
  server --master-port=20003 --web-port=10003
```

