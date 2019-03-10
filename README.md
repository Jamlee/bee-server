bee-server
========
[![CodeFactor](https://www.codefactor.io/repository/github/jamlee/bee-server/badge)](https://www.codefactor.io/repository/github/jamlee/bee-server)
[![Build Status](https://travis-ci.com/Jamlee/bee-server.svg?branch=master)](https://travis-ci.com/Jamlee/bee-server)

bee-server is a simple and stable game server cluster.


---

### Building

building need docker running on the host.

```
# check golang code quality and build to bin
make
```

### Starting the cluster

it is less 3 nodes to be required for **HA**.

```
./start -c start.conf
```