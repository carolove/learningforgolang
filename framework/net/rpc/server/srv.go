package server

import (
	"github.com/cuigh/auxo/config"
	"github.com/cuigh/auxo/net/rpc"

	_ "github.com/carolove/learningforgolang/framework/net/rpc/registry"
	"github.com/cuigh/auxo/data"
	"github.com/cuigh/auxo/net/transport"
)

func Listen(addrs ...transport.Address) (*rpc.Server, error) {
	opts := rpc.ServerOptions{
		Address: addrs,
	}
	opts.Registry = struct {
		Name    string   `json:"name" yaml:"name"`
		Options data.Map `json:"options" yaml:"options"`
	}{}
	opts.Registry.Name = config.GetString("global.registry.store.backend")
	return rpc.NewServer(opts)
}
