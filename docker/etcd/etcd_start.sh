#!/bin/bash

ETCD_VERSION=v3.2.15

TOKEN=my-etcd-token

CLUSTER_STATE=new

NAME_1=etcd-node-0

HOST_1=127.0.0.1

CLUSTER=${NAME_1}=http://${HOST_1}:2380

# For node 1

THIS_NAME=${NAME_1}

THIS_IP=${HOST_1}

docker pull xieyanze/etcd3:latest

docker run -d -p 2379:2379 -p 2380:2380 -p 4001:4001 -p 7001:7001 --name etcd-v3 quay.io/coreos/etcd:${ETCD_VERSION} \
	/usr/local/bin/etcd
