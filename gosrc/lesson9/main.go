package main

import (
	"fmt"

	"github.com/abronan/valkeyrie/store"
	"github.com/cuigh/auxo/data"
	"github.com/cuigh/auxo/net/rpc/registry"
	"github.com/cuigh/auxo/net/rpc/registry/valkeyrie"
	"github.com/cuigh/auxo/net/transport"
)

func main() {
	// register etcd2 backend
	_backend := store.ETCD
	ret := &valkeyrie.Builder{Backend: _backend}
	registry.Register(ret)
	s := registry.Server{
		Name:    "demo",
		Version: "1.0.0",
		Addresses: []transport.Address{
			{URL: "localhost:2379"},
		},
	}
	opts := data.Map{
		"address": "localhost:2379",
	}
	r, err := registry.Get(ret.Name()).Build(s, opts)
	if err != nil {
		fmt.Println(err)
	}
	r.Online()
}
